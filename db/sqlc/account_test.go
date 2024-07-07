package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tommydebisi/go-simple-bank/utils"
)

// creates a new Account
func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner: utils.RandomOwner(),
		Balance: utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	// check the values of arg is same as created
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	// check that auto generated fields
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)
	
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.Equal(t, acc1.Owner, acc2.Owner)
	
	// check if two times are within duration of each other
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	acc1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID: acc1.ID,
		Balance: utils.RandomMoney(),
	}
	acc2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)
	
	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, arg.Balance, acc2.Balance)
}

func TestDeleteAccount(t *testing.T) {
	acc1 := createRandomAccount(t)

	// delete account created
	err := testQueries.DeleteAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	// get the Id to check
	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, acc2)
}

func TestListAccounts(t *testing.T) {
	// create multiple accounts
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit: 5,
		Offset: 1,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	require.NotZero(t, len(accounts))
	// check if all accounts in slice are not empty
	for _, acccount := range accounts{
		require.NotEmpty(t, acccount)
	}

}