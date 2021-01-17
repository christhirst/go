package db

import (
	"context"
	"database/sql"
	"go-postgres/db/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomDemand(t *testing.T) Demand {
	arg := CreateDemandParams{
		Title:       util.RandomUser(),
		ID:          util.RandomString(4),
		Owner:       sql.NullString{String: util.RandomString(4), Valid: true},
		Description: sql.NullString{String: util.RandomString(4), Valid: true},
	}
	demand, err := testQueries.CreateDemand(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, demand)

	require.Equal(t, arg.Title, demand.Title)
	require.Equal(t, arg.ID, demand.ID)

	require.NotZero(t, demand.ID)

	return demand

}
func TestCreateDemand(t *testing.T) {
	createRandomDemand(t)
}

func TestGetDemand(t *testing.T) {
	demand1 := createRandomDemand(t)
	demand2, err := testQueries.GetDemand(context.Background(), demand1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, demand2)

	require.Equal(t, demand1.Title, demand2.Title)
	require.Equal(t, demand1.ID, demand2.ID)

}

func TestUpdateDemand(t *testing.T) {
	demand1 := createRandomDemand(t)

	arg := UpdateDemandParams{
		Title:       demand1.Title,
		Owner:       demand1.Owner,
		Description: demand1.Description,
	}

	demand2, err := testQueries.UpdateDemand(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, demand2)

	require.Equal(t, demand1.Title, demand2.Title)
	require.Equal(t, arg.Owner, demand2.Owner)
}

func TestDeleteDemand(t *testing.T) {
	demand1 := createRandomDemand(t)
	err := testQueries.DeleteDemand(context.Background(), demand1.ID)

	require.NoError(t, err)

	demand2, err := testQueries.GetDemand(context.Background(), demand1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, demand2)
}

func TestListDemand(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomDemand(t)
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
