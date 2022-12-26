package usecase_test

import (
	"testing"

	"github.com/kpaya/wallet-go/internal/entity"
	usecase "github.com/kpaya/wallet-go/internal/usecase/create_account"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (uc *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := uc.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (uc *AccountGatewayMock) Save(account *entity.Account) error {
	args := uc.Called(account)
	return args.Error(0)
}

type ClientGatewayMock struct {
	mock.Mock
}

func (uc *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := uc.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (uc *ClientGatewayMock) Save(client *entity.Client) error {
	args := uc.Called(client)
	return args.Error(0)
}

func TestCreateAccount(t *testing.T) {
	client, err := entity.NewClient("Pipoca", "pipoca@gmail.com")
	require.Nil(t, err)
	clientMock := &ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc, err := usecase.NewCreateAccount(accountMock, clientMock)
	require.Nil(t, err)

	output, err := uc.Execute(usecase.InputCreateAccountDTO{
		ClientID: client.ID,
	})
	require.Nil(t, err)
	require.NotNil(t, output.ID)

	clientMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "Save", 1)

}
