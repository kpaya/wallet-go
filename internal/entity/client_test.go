package entity_test

import (
	"testing"

	"github.com/kpaya/wallet-go/internal/entity"
	"github.com/stretchr/testify/require"
)

func TestCreateANewClient(t *testing.T) {
	_, err := entity.NewClient("Pipoca", "pipoca@gmail.com")
	require.Nil(t, err)
}

func TestCreateANewClientError(t *testing.T) {
	_, err := entity.NewClient("", "")
	require.Nil(t, err)
}
