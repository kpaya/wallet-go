package usecase

import (
	"time"

	"github.com/kpaya/wallet-go/internal/entity"
	"github.com/kpaya/wallet-go/internal/gateway"
)

type InputCreateClientDTO struct {
	Name  string
	Email string
}

type OutputCreateClientDTO struct {
	ID        string    `json:"id" validate:"uuid,required"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"email,required"`
	CreatedAt time.Time `json:"created_at" validate:"-"`
	UpdatedAt time.Time `json:"updated_at" validate:"-"`
}

type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(gateway gateway.ClientGateway) (*CreateClientUseCase, error) {
	return &CreateClientUseCase{
		ClientGateway: gateway,
	}, nil
}

func (c *CreateClientUseCase) Execute(input *InputCreateClientDTO) (*OutputCreateClientDTO, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}
	err = c.ClientGateway.Save(client)
	if err != nil {
		return nil, err
	}
	return &OutputCreateClientDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil

}
