package db

import (
	"context"
	"testing"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDb)

	_ = store.execTx(context.Background(), func(q *Queries) error {

	})
}
