package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDb)

	// create transfers to send
	account := createRandomAccount(t)
	account2 := createRandomAccount(t)
	amount := int64(10)
	nTimes := 5

	// create channel to store err and response of goroutines
	errs := make(chan error)
	responses := make(chan TransferTxResponse)

	// run series of transfer transactions concurrently
	for i := 0; i < nTimes; i++ {
		go func() {
			res, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: account.ID,
				ToAccountId:   account2.ID,
				Amount:        int64(amount),
			})

			errs <- err
			responses <- res
		}()
	}

	// go through nTimes of tx performed and error check
	for i := 0; i < nTimes; i++ {
		err := <-errs
		require.NoError(t, err)

		res := <-responses
		require.NotEmpty(t, res)

		// check transfers
		transfers := res.Transfer
		require.NotEmpty(t, transfers)

		// check if the transfer record was created properly
		require.Equal(t, account.ID, transfers.FromAccountID)
		require.Equal(t, account2.ID, transfers.ToAccountID)
		require.Equal(t, amount, transfers.Amount)
		require.NotZero(t, transfers.CreatedAt)
		require.NotZero(t, transfers.ID)

		// check if the transfer is in the db
		_, tfErr := store.GetTransfer(context.Background(), transfers.ID)
		require.NoError(t, tfErr)

		// check entry from
		entryFrom := res.EntryFrom
		require.NotEmpty(t, entryFrom)
		require.Equal(t, account.ID, entryFrom.AccountID)
		require.Equal(t, entryFrom.Amount, -amount)
		require.NotEmpty(t, entryFrom.CreatedAt)
		require.NotEmpty(t, entryFrom.ID)

		// check if entry from got created in the db
		_, efErr := store.GetEntry(context.Background(), entryFrom.ID)
		require.NoError(t, efErr)

		// check entry to
		entryTo := res.EntryTo
		require.NotEmpty(t, entryTo)
		require.Equal(t, account2.ID, entryTo.AccountID)
		require.Equal(t, entryTo.Amount, amount)
		require.NotEmpty(t, entryTo.CreatedAt)
		require.NotEmpty(t, entryTo.ID)

		// check if entry to got created in the db
		_, etErr := store.GetEntry(context.Background(), entryTo.ID)
		require.NoError(t, etErr)

		// check account from
		accountFrom := res.AccountFrom
		require.NotEmpty(t, accountFrom)
		require.Equal(t, accountFrom.ID, account.ID)

		// check account to
		accountTo := res.AccountTo
		require.NotEmpty(t, accountTo)
		require.Equal(t, accountTo.ID, account.ID)

		// check amount transfered is the same received
		diff1 := account.Balance - accountFrom.Balance
		diff2 := accountTo.Balance - account2.Balance

		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)         // the amount transfered should be positive
		require.True(t, diff1%amount == 0) // depending on ntimes the remainder should be nil

		k := int(diff1 / amount)
		require.True(t, k > 1 && k <= nTimes) // k should be the current index plus 1
	}
}
