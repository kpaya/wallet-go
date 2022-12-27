package gateway

import "github.com/kpaya/wallet-go/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
