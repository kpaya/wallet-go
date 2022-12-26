package gateway

import "github.com/kpaya/wallet-go/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
