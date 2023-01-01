package database_test

import (
	"database/sql"
	"testing"

	"github.com/kpaya/wallet-go/internal/database"
	"github.com/kpaya/wallet-go/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	transactionDb *database.TransactionDB
	clientFrom    *entity.Client
	clientTo      *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	db.Exec(`CREATE TABLE clients (id varchar(36), name varchar(1000), email varchar(1000), created_at date)`)
	s.Nil(err)
	db.Exec(`CREATE TABLE accounts (id varchar(36), client_id varchar(36), balance int, created_at date)`)
	s.Nil(err)
	db.Exec(`CREATE TABLE transactions (id varchar(36), account_id_from varchar(36), account_id_to varchar(36), amount int, created_at date)`)
	s.Nil(err)
	s.clientFrom, err = entity.NewClient("clientFrom", "clientFrom@gmail.com")
	s.Nil(err)
	s.clientTo, err = entity.NewClient("clientTo", "clientTo@gmail.com")
	s.Nil(err)
	s.accountFrom, err = entity.NewAccount(s.clientFrom)
	s.accountFrom.Balance = 1000.0
	s.Nil(err)
	s.accountTo, err = entity.NewAccount(s.clientTo)
	s.accountTo.Balance = 1000.0
	s.Nil(err)

	s.transactionDb = database.NewTransactionDb(db)
}

func (p *TransactionDBTestSuite) TearDownSuite(t *testing.T) {
	defer p.db.Close()
	p.db.Exec("DROP TABLE clients;")
	p.db.Exec("DROP TABLE accounts;")
	p.db.Exec("DROP TABLE addresses;")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 1000)
	s.Nil(err)

	err = s.transactionDb.Create(transaction)
	s.Nil(err)
}
