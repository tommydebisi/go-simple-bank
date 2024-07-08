package db

import (
	"context"
	"testing"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDb)

	// create transfers to send
	transfer := createRandTransfer(t, createRandomAccount(t).ID, createRandomAccount(t).ID)
	transfer2 := createRandTransfer(t, createRandomAccount(t).ID, createRandomAccount(t).ID)

	// create channel
	// run series of transfer transactions concurrently
	for i := 0; i < 5; i++ {
		go func() {
			store.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: transfer.ID,
				ToAccountId:   transfer2.ID,
				Amount:        10,
			})
		}()
	}
}
