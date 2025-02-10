package pgtest

import (
	"os"

	"github.com/Maksim646/ozon_shortner/internal/database/postgresql"
	"github.com/heetch/sqalx"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	postgresURI  string
	migrationDir string
	migrator     *postgresql.Migrator
	db           sqalx.Node
	sqlxDB       *sqlx.DB
}

func (s *Suite) DB() sqalx.Node { return s.db }

func (s *Suite) SetupSuite(relativePath string) {
	s.migrationDir = relativePath + "/database/postgresql/migrations"
	s.postgresURI = os.Getenv("TEST_POSTGRES_URI")
	if s.postgresURI == "" {
		s.postgresURI = "postgres://postgres:ozon_shortner@localhost:5448/ozon_shortner_db_test?sslmode=disable"
	}

	sqlxDB, err := sqlx.Connect("postgres", s.postgresURI)
	if err != nil {
		s.FailNowf("cannot open db connection; forgot to run make db?", "cannot open db connection to testing db(%s): %v", s.postgresURI, err)
	}
	s.sqlxDB = sqlxDB

	s.migrator = postgresql.NewMigrator(s.postgresURI, s.migrationDir)

	err = s.migrator.Apply()
	if err != nil {
		s.FailNowf("cannot apply migrations", "cannot apply migrations: %v", err)
		return
	}
}

func (s *Suite) SetupTest() {
	dbNode, err := sqalx.New(s.sqlxDB)
	if err != nil {
		s.FailNowf("cannot create sqlx node", "cannot create sqlx node: %v", err)
	}

	db, err := dbNode.Beginx()
	if err != nil {
		s.FailNowf("cannot start transaction", "cannot start transaction: %v", err)
	}

	s.db = db
}

func (s *Suite) TearDownTest() {
	s.Nil(s.db.Rollback())
}

func (s *Suite) TearDownSuite() {
	if err := s.migrator.Revert(); err != nil {
		s.FailNowf("cannot revert migrations", "cannot revert migrations: %v", err)
	}
	s.sqlxDB.Close()
	os.Unsetenv("MIGRATIONS_APPLIED")
}
