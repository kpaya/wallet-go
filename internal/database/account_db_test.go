package database_test

import (
	"database/sql"
	"testing"

	"github.com/kpaya/wallet-go/internal/database"
	"github.com/kpaya/wallet-go/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	client    *entity.Client
	accountDb *database.AccountDb
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec(`CREATE TABLE clients (id varchar(36), name varchar(1000), email varchar(1000), created_at date, updated_at date)`)
	db.Exec(`CREATE TABLE accounts (id varchar(36), client_id varchar(36), balance int, created_at date, updated_at date)`)
	s.client, _ = entity.NewClient("Pitomba", "pipoca@gmail.com")
	db.Exec(`INSERT INTO clients (id, name, email, created_at, updated_at) VALUES (?,?,?,?,?)`, s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt, s.client.UpdatedAt)
	s.NotNil(s.client.ID)
	s.accountDb = database.NewAccountDb(db)
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDbTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {

	account, err := entity.NewAccount(s.client)
	s.Nil(err)
	err = s.accountDb.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindById() {

	account, err := entity.NewAccount(s.client)
	s.Nil(err)
	err = s.accountDb.Save(account)
	s.Nil(err)

	accountDb, err := s.accountDb.FindById(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDb.ID)
	s.Equal(account.Client.ID, accountDb.Client.ID)
}
