package entity_test

import (
	"testing"

	"github.com/kpaya/wallet-go/internal/entity"
	"github.com/stretchr/testify/require"
)

func TestCreateANewTransaction(t *testing.T) {
	client1, err := entity.NewClient("Lucas De Sá", "lucas@gmail.com")
	require.Nil(t, err)
	require.NotNil(t, client1.ID)
	client2, err := entity.NewClient("Luan de Sá", "luan@gmail.com")
	require.Nil(t, err)
	require.NotNil(t, client2.ID)

	account1, err := entity.NewAccount(client1)
	require.Nil(t, err)
	require.NotNil(t, account1.Client.ID)

	account2, err := entity.NewAccount(client2)
	require.Nil(t, err)
	require.NotNil(t, account2.Client.ID)

	account1.Credit(1500.00)

	transaction, err := entity.NewTransaction(account1, account2, 450.0)
	require.Nil(t, err)
	require.NotNil(t, transaction.ID)

	require.Equal(t, 1050.0, account1.Balance)
	require.Equal(t, 450.0, account2.Balance)
}
