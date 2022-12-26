package usecase_test

import (
	"testing"

	"github.com/kpaya/wallet-go/internal/entity"
	usecase "github.com/kpaya/wallet-go/internal/usecase/create_client"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func TestCreateClientUseCase_Execute(t *testing.T) {

	m := &ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)

	uc, err := usecase.NewCreateClientUseCase(m)
	require.Nil(t, err)

	output, err := uc.Execute(&usecase.InputCreateClientDTO{
		Name:  "Pipoca Doce",
		Email: "pipoca@gmail.com",
	})

	require.Nil(t, err)
	require.Equal(t, "Pipoca Doce", output.Name)
	require.Equal(t, "pipoca@gmail.com", output.Email)

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
