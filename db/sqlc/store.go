package db

import (
	"context"
	"database/sql"
)

// Provides all functions for queries and transactions
type Store struct {
	// extension of queries, called a composition (like inheritance)
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries: New(db),
	}
}

func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
}