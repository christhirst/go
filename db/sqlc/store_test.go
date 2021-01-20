package db

import (
	"context"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestTrasferTx(t *testing.T) {

	store := NewStore(testDB)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	demand1 := createRandomDemand(t)

	n := n

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.UserID,
				ToAccountID:   account2.UserID,
				DemandID:      demand1.ID,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs

		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)

		transfer := result.DemandTransfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.UserID, transfer.FromAccountID)


		require.NoError(t, err)
	}

}
