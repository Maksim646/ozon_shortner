package pgtest

import (
	"fmt"
	"log"
	"os"
	"strings"

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

func (s *Suite) connectToPostgres(postgresURI string) (*sqlx.DB, error) {
	sqlxDB, err := sqlx.Connect("postgres", postgresURI)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres database (%s): %w", postgresURI, err)
	}
	log.Printf("Successfully connected to postgres database: %s", postgresURI) // Используем логирование

	return sqlxDB, nil
}

func (s *Suite) DropDB() error {
	postgresURI := strings.Replace(s.postgresURI, "/ozon_shortner_db_test", "/postgres", 1)
	postgresDB, err := s.connectToPostgres(postgresURI)
	if err != nil {
		return fmt.Errorf("failed to connect to 'postgres' db for dropping test db: %w", err)
	}
	defer postgresDB.Close()

	_, err = postgresDB.Exec(`
		SELECT pg_terminate_backend(pg_stat_activity.pid)
		FROM pg_stat_activity
		WHERE pg_stat_activity.datname = 'ozon_shortner_db_test'
		  AND pg_stat_activity.pid <> pg_backend_pid();
	`)
	if err != nil {
		return fmt.Errorf("failed to terminate other database sessions: %w", err)
	}

	_, err = postgresDB.Exec("DROP DATABASE IF EXISTS ozon_shortner_db_test;")
	if err != nil {
		return fmt.Errorf("failed to drop database 'ozon_shortner_db_test': %w", err)
	}

	log.Println("Successfully dropped database: ozon_shortner_db_test") // Используем логирование
	return nil
}

func (s *Suite) CreateDB() error {
	postgresURI := strings.Replace(s.postgresURI, "/ozon_shortner_db_test", "/postgres", 1)
	postgresDB, err := s.connectToPostgres(postgresURI)
	if err != nil {
		return fmt.Errorf("failed to connect to 'postgres' db for creating test db: %w", err)
	}
	defer postgresDB.Close()

	_, err = postgresDB.Exec("CREATE DATABASE ozon_shortner_db_test;")
	if err != nil {
		return fmt.Errorf("failed to create database 'ozon_shortner_db_test': %w", err)
	}

	log.Println("Successfully created database: ozon_shortner_db_test") // Используем логирование
	return nil
}

func (s *Suite) SetupSuite(relativePath string) {
	s.migrationDir = relativePath + "/database/postgresql/migrations"
	s.postgresURI = os.Getenv("TEST_POSTGRES_URI")
	if s.postgresURI == "" {
		s.postgresURI = "postgres://postgres:ozon_shortner@localhost:5448/ozon_shortner_db_test?sslmode=disable"
	}

	sqlxDB, err := s.connectToPostgres(s.postgresURI)
	if err != nil {
		s.FailNowf("cannot open db connection; forgot to run make db?", "cannot open db connection to testing db(%s): %v", s.postgresURI, err)
	}
	s.sqlxDB = sqlxDB

	if s.sqlxDB == nil {
		s.FailNowf("sqlxDB is nil", "sqlxDB is nil, cannot proceed with DropDB and CreateDB")
		return
	}

	err = s.DropDB()
	if err != nil {
		s.FailNowf("cannot drop db", "failed to drop testing db: %v", err) // Улучшенное сообщение об ошибке
	}
	err = s.CreateDB()
	if err != nil {
		s.FailNowf("cannot create db", "failed to create testing db(%s): %v", s.postgresURI, err) // Улучшенное сообщение об ошибке
	}

	sqlxDB, err = s.connectToPostgres(s.postgresURI)
	if err != nil {
		s.FailNowf("cannot open db connection after creating db", "cannot open db connection to testing db(%s): %v", s.postgresURI, err)
	}
	s.sqlxDB = sqlxDB

	if s.sqlxDB == nil {
		s.FailNowf("sqlxDB is nil", "sqlxDB is nil, cannot proceed with DropDB and CreateDB")
		return
	}

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
	err := s.db.Rollback()
	if err != nil {
		s.FailNowf("cannot rollback transaction", "cannot rollback transaction: %v", err)
	}
}

func (s *Suite) TearDownSuite() {
	if err := s.migrator.Revert(); err != nil {
		s.FailNowf("cannot revert migrations", "cannot revert migrations: %v", err)
	}
	s.sqlxDB.Close()
	os.Unsetenv("MIGRATIONS_APPLIED")
}
