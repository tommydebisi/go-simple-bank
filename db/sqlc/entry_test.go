package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tommydebisi/go-simple-bank/utils"
)

func createRandEntry(t *testing.T, accountId, amount int64) Entry {
	arg := CreateEntryParams{
		AccountID: accountId,
		Amount: amount,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandEntry(t, createRandomAccount(t).ID, utils.RandomMoney())
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandEntry(t, createRandomAccount(t).ID, utils.RandomMoney())

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	
	require.NotZero(t, entry2.ID)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	accountId := createRandomAccount(t).ID
	
	for i := 0; i < 10; i++ {
		createRandEntry(t, accountId, utils.RandomMoney())
	}

	arg := ListEntriesParams{
		AccountID: accountId,
		Limit: 5,
		Offset: 3,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}