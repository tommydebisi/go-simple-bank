package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tommydebisi/go-simple-bank/utils"
)

func createRandTransfer(t *testing.T, from, to int64) Transfer {
	arg := CreateTransferParams{
		FromAccountID: from,
		ToAccountID: to,
		Amount: utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandTransfer(t, createRandomAccount(t).ID, createRandomAccount(t).ID)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandTransfer(t, createRandomAccount(t).ID, createRandomAccount(t).ID)
	
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	arg := ListTransfersParams{
		FromAccountID: createRandomAccount(t).ID,
		ToAccountID: createRandomAccount(t).ID,
		Limit: 5,
		Offset: 2,
	}
	
	for i := 0; i < 10; i++ {
		createRandTransfer(t, arg.FromAccountID, arg.FromAccountID)
	}
	
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}