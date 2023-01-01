package database_test

import (
	"database/sql"
	"testing"

	"github.com/kpaya/wallet-go/internal/database"
	"github.com/kpaya/wallet-go/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDb *database.ClientDb
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec(`CREATE TABLE clients (id varchar(36), name varchar(1000), email varchar(1000), created_at date, updated_at date)`)
	s.clientDb = database.NewClientDb(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDbTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("Lucas", "lucas@gmail.com")
	s.clientDb.Save(client)

	clientDb, err := s.clientDb.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDb.ID)
	s.Equal(client.Name, clientDb.Name)
	s.Equal(client.Email, clientDb.Email)
}
