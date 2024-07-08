package db

import (
	"testing"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDb)

	transfer := createRandTransfer(t)
}
