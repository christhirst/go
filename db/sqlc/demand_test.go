package db

import (
	"context"
	"database/sql"
	"go-postgres/db/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomDemand(t *testing.T, account Account) Demand {
	arg := CreateDemandParams{
		Title:       util.RandomUser(),
		AccountID:   account.UserID,
		Description: sql.NullString{String: util.RandomString(4), Valid: true},
	}
	demand, err := testQueries.CreateDemand(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, demand)

	require.Equal(t, arg.Title, demand.Title)

	require.NotZero(t, demand.ID)

	return demand

}
func TestCreateDemand(t *testing.T) {
	account := createRandomAccount(t)
	createRandomDemand(t, account)
}

func TestGetDemand(t *testing.T) {
	account := createRandomAccount(t)
	demand1 := createRandomDemand(t, account)

	demand2, err := testQueries.GetDemand(context.Background(), demand1.Title)

	require.NoError(t, err)
	require.NotEmpty(t, demand2)

	require.Equal(t, demand1.Title, demand2.Title)
	require.Equal(t, demand1.ID, demand2.ID)

}

func TestUpdateDemand(t *testing.T) {
	account := createRandomAccount(t)
	demand1 := createRandomDemand(t, account)

	arg := UpdateDemandParams{
		AccountID:   demand1.AccountID,
		Description: demand1.Description,
	}

	demand2, err := testQueries.UpdateDemand(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, demand2)

	require.Equal(t, demand1.Title, demand2.Title)
	require.Equal(t, arg.AccountID, demand2.AccountID)
}

func TestDeleteDemand(t *testing.T) {
	account := createRandomAccount(t)
	demand1 := createRandomDemand(t, account)
	err := testQueries.DeleteDemand(context.Background(), demand1.ID)

	require.NoError(t, err)

	demand2, err := testQueries.GetDemand(context.Background(), demand1.Title)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, demand2)
}

func TestListDemand(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomDemand(t, account)
	}

	arg := ListDemandsParams{
		Limit:  5,
		Offset: 5,
	}
	demands, err := testQueries.ListDemands(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, demands, 5)

	for _, demand := range demands {
		require.NotEmpty(t, demand)
	}
}
