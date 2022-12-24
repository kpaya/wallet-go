package entity_test

import (
	"testing"

	"github.com/kpaya/wallet-go/internal/entity"
	"github.com/stretchr/testify/require"
)

func TestCreateANewAccount(t *testing.T) {
	user, _ := entity.NewClient("Lucas de Sá", "lucas@gmail.com")
	require.NotNil(t, user.ID)

	account, err := entity.NewAccount(user)
	require.Nil(t, err)
	require.Equal(t, account.Balance, 0.0)

	err = account.Credit(23.22)
	require.Equal(t, 23.22, account.Balance)
	require.Nil(t, err)

	err = account.Debit(10)
	require.Nil(t, err)

	require.Equal(t, float64(13.22), account.Balance)
}
