package usecase

import (
	"github.com/kpaya/wallet-go/internal/entity"
	"github.com/kpaya/wallet-go/internal/gateway"
)

type InputCreateTransactionDTO struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type OutputCreateTransactionDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGetway      gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGetway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGetway:      accountGetway,
	}
}

func (t *CreateTransactionUseCase) Execute(input *InputCreateTransactionDTO) (*OutputCreateTransactionDTO, error) {
	accountFrom, err := t.AccountGetway.FindById(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}
	accountTo, err := t.AccountGetway.FindById(input.AccountIDTo)
	if err != nil {
		return nil, err
	}
	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = t.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	return &OutputCreateTransactionDTO{
		ID: transaction.ID,
	}, nil
}
