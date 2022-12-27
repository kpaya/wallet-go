package usecase_test

import (
	"testing"

	"github.com/kpaya/wallet-go/internal/entity"
	usecase "github.com/kpaya/wallet-go/internal/usecase/create_transaction"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

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

func TestCreateTransaction_Execute(t *testing.T) {
	client1, _ := entity.NewClient("Pipoca Doce 1", "pipoca1@gmail.com")
	client2, _ := entity.NewClient("Pipoca Doce 2", "pipoca2@gmail.com")

	account1, _ := entity.NewAccount(client1)
	account1.Credit(500)
	account2, _ := entity.NewAccount(client2)
	account2.Credit(500)

	mockTransactionGateway := TransactionGatewayMock{}
	mockTransactionGateway.On("Create", mock.Anything).Return(nil)
	mockAccountGateway := AccountGatewayMock{}
	mockAccountGateway.On("FindById", account1.ID).Return(account1, nil)
	mockAccountGateway.On("FindById", account2.ID).Return(account2, nil)

	uc := usecase.NewCreateTransactionUseCase(&mockTransactionGateway, &mockAccountGateway)

	output, err := uc.Execute(&usecase.InputCreateTransactionDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        500,
	})

	require.Nil(t, err)
	require.NotNil(t, output.ID)
	mockTransactionGateway.AssertExpectations(t)
	mockTransactionGateway.AssertNumberOfCalls(t, "Create", 1)
	mockAccountGateway.AssertExpectations(t)
	mockAccountGateway.AssertNumberOfCalls(t, "FindById", 2)
}
