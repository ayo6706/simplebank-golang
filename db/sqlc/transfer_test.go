package db

import (
	"context"
	"testing"

	"github.com/ayo6706/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T, account1 Account, account2 Account) Transfer {
	arg := CreateTransferParams{
		ToAccountID:   account1.ID,
		FromAccountID: account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.Amount, transfer.Amount)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)

	return transfer

}

func TestCreateTransfer(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	CreateRandomTransfer(t, account1, account2)
}