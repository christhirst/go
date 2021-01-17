package db

import (
	"context"
	"database/sql"
	"go-postgres/db/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		UserID:   util.RandomString(4),
		Username: util.RandomUser(),
		Password: util.RandomString(4),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Username, account.Username)
	require.Equal(t, arg.Password, account.Password)

	require.NotZero(t, account.Username)

	return account

}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.Username, account2.Username)
	require.Equal(t, account1.Password, account2.Password)

}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		Username: account1.Username,
		Password: util.RandomString(4),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.Username, account2.Username)
	require.Equal(t, arg.Password, account2.Password)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.Username)

	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.Username)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
