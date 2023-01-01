package database

import (
	"database/sql"

	"github.com/kpaya/wallet-go/internal/entity"
)

type ClientDb struct {
	DB *sql.DB
}

func NewClientDb(db *sql.DB) *ClientDb {
	return &ClientDb{DB: db}
}

func (c *ClientDb) Get(id string) (*entity.Client, error) {
	var client entity.Client
	row := c.DB.QueryRow("SELECT id, name, email, created_at FROM clients WHERE id = ?", id)
	if err := row.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {
		return nil, err
	}
	return &client, nil
}

func (c *ClientDb) Save(client *entity.Client) error {
	stmt, err := c.DB.Prepare("INSERT INTO clients (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(client.ID, client.Name, client.Email, client.CreatedAt, client.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
