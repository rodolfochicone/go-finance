package db

import (
	"context"
	"github.com/cip8/autoname"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	category := createRandomCategory(t)
	arg := CreateAccountParams{
		UserID:      category.UserID,
		Title:       autoname.Generate(""),
		Type:        "debit",
		Description: autoname.Generate(" "),
		CategoryID:  category.ID,
		Date:        time.Now(),
		Value:       10,
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.UserID, account.UserID)
	require.Equal(t, arg.CategoryID, account.CategoryID)
	require.Equal(t, arg.Value, account.Value)
	require.Equal(t, arg.Title, account.Title)
	require.Equal(t, arg.Type, account.Type)
	require.Equal(t, arg.Description, account.Description)

	require.NotEmpty(t, account.CreatedAt)
	require.NotEmpty(t, account.Date)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccountById(t *testing.T) {
	account1 := createRandomAccount(t)

	account2, err := testQueries.GetAccountsById(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.UserID, account2.UserID)
	require.Equal(t, account1.CategoryID, account2.CategoryID)
	require.Equal(t, account1.Value, account2.Value)
	require.Equal(t, account1.Title, account2.Title)
	require.Equal(t, account1.Type, account2.Type)
	require.Equal(t, account1.Description, account2.Description)

	require.NotEmpty(t, account2.CreatedAt)
	require.NotEmpty(t, account2.Date)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccountsById(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, account2)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:          account1.ID,
		Title:       autoname.Generate(""),
		Description: autoname.Generate(" "),
		Value:       15,
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, arg.Title, account2.Title)
	require.Equal(t, arg.Description, account2.Description)
	require.Equal(t, arg.Value, account2.Value)
	require.Equal(t, account1.CreatedAt, account2.CreatedAt)
}

func TestListAccounts(t *testing.T) {
	lastAccount := createRandomAccount(t)

	arg := GetAccountsParams{
		UserID:      lastAccount.UserID,
		Type:        lastAccount.Type,
		Title:       lastAccount.Title,
		Description: lastAccount.Description,
		CategoryID:  lastAccount.CategoryID,
		Date:        lastAccount.Date,
	}

	accounts, err := testQueries.GetAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 1)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.Equal(t, lastAccount.ID, account.ID)
		require.Equal(t, lastAccount.Title, account.Title)
		require.Equal(t, lastAccount.UserID, account.UserID)
		require.Equal(t, lastAccount.Description, account.Description)
		require.Equal(t, lastAccount.Value, account.Value)
		require.NotEmpty(t, lastAccount.CreatedAt)
		require.NotEmpty(t, lastAccount.Date)
		log.Println("Category Title:", account.CategoryTitle)
	}
}

func TestSumValues(t *testing.T) {
	lastAccount := createRandomAccount(t)

	arg := GetAccountsReportsParams{
		UserID: lastAccount.UserID,
		Type:   lastAccount.Type,
	}

	sumValue, err := testQueries.GetAccountsReports(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sumValue)
}

func TestCountReports(t *testing.T) {
	lastAccount := createRandomAccount(t)

	arg := GetAccountsGraphParams{
		UserID: lastAccount.UserID,
		Type:   lastAccount.Type,
	}

	count, err := testQueries.GetAccountsGraph(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, count)
}
