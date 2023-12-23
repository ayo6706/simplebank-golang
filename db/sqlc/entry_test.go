package db

import (
	"context"
	"testing"

	"github.com/ayo6706/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	account := CreateRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, account.ID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.CreatedAt)

}
