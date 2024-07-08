package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Provides all functions for queries and transactions
type Store struct {
	// extension of queries, called a composition (like inheritance)
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// executes a transaction passed
func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queries := New(tx)
	err = fn(queries)
	if err != nil {
		// rollback if error occurs when performing queries
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			return fmt.Errorf("Error rolling back transaction: %v", rollBackErr)
		}
		return err
	}

	// commit changes
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResponse struct {
	Transfer    Transfer `json:"transfer"`
	AccountFrom Account  `json:"account_from"`
	AccountTo   Account  `json:"account_to"`
	EntryFrom   Entry    `json:"entry_from"`
	EntryTo     Entry    `json:"entry_to"`
}

// transfers money from one account to the other ...
// creates a transfer record, then two entry records for each account, updates both accounts using the entry records
func (s *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResponse, error) {
	var response TransferTxResponse

	// start executing a transaction
	_ = s.execTx(ctx, func(q *Queries) error {
		var err error

		// create a new transfer record
		response.Transfer, err = s.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID:   arg.ToAccountId,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// create entry for account sending
		response.EntryFrom, err = s.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountId,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		// create entry for account receiving
		response.EntryFrom, err = s.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountId,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// TODO: update account to be done

		return nil
	})

	return response, nil
}
