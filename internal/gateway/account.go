package gateway

import "github.com/kpaya/wallet-go/internal/entity"

type AccountGateway interface {
	FindById(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}
