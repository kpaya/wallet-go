package database

import (
	"database/sql"

	"github.com/kpaya/wallet-go/internal/entity"
)

type TransactionDB struct {
	Db *sql.DB
}

func NewTransactionDb(db *sql.DB) *TransactionDB {
	return &TransactionDB{
		Db: db,
	}
}

func (t *TransactionDB) Create(transaction *entity.Transaction) error {
	stmt, err := t.Db.Prepare("INSERT INTO transactions (id, account_id_from, account_id_to, amount, created_at) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(transaction.ID, transaction.AccountFrom.ID, transaction.AccountTo.ID, transaction.Amount, transaction.CreatedAt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
