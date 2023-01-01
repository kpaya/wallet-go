package database

import (
	"database/sql"

	"github.com/kpaya/wallet-go/internal/entity"
)

type AccountDb struct {
	Db *sql.DB
}

func NewAccountDb(db *sql.DB) *AccountDb {
	return &AccountDb{
		Db: db,
	}
}

func (a *AccountDb) FindById(id string) (*entity.Account, error) {
	var account entity.Account
	var client entity.Client
	account.Client = &client
	row := a.Db.QueryRow("SELECT ac.id, ac.client_id, ac.balance, cl.name, cl.email FROM accounts ac INNER JOIN clients cl ON cl.id = ac.client_id WHERE ac.id = ?", id)
	if err := row.Scan(&account.ID, &client.ID, &account.Balance, &client.Name, &client.Email); err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *AccountDb) Save(account *entity.Account) error {
	stmt, err := a.Db.Prepare("INSERT INTO accounts (id, client_id, balance, created_at, updated_at) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(account.ID, account.Client.ID, account.Balance, account.CreatedAt, account.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
