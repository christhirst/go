package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transaction
type Store struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID string `json:"from_account_id"`
	ToAccountID   string `json:"to_account_id"`
	DemandID      string `json:"demand"`
}

type TransferTxResult struct {
	DemandTransfer DemandTransfer `json:"demandTransfer"`
	FromAccount    Account        `json:"from_account"`
	ToAccount      Account        `json:"to_account"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {

	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.DemandTransfer, err = q.CreateDemand_transfer(ctx, CreateDemand_transferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			DemandID:      arg.DemandID,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
