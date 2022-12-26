package usecase

import (
	"github.com/kpaya/wallet-go/internal/entity"
	"github.com/kpaya/wallet-go/internal/gateway"
)

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

type InputCreateAccountDTO struct {
	ClientID string
}

type OutputCreateAccountDTO struct {
	ID string
}

func NewCreateAccount(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) (*CreateAccountUseCase, error) {
	return &CreateAccountUseCase{
		AccountGateway: accountGateway,
		ClientGateway:  clientGateway,
	}, nil
}

func (c *CreateAccountUseCase) Execute(input InputCreateAccountDTO) (*OutputCreateAccountDTO, error) {
	client, err := c.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}
	account, err := entity.NewAccount(client)
	if err != nil {
		return nil, err
	}
	err = c.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}
	return &OutputCreateAccountDTO{ID: account.ID}, nil
}
