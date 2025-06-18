package database

import (
	"errors"
	"fmt"
	"os"

	migrator "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // import for migration file
	"github.com/jmoiron/sqlx"
	_ "github.com/mattes/migrate/source/file" // import for migration file
)

var (
	NoteDB *sqlx.DB
)

type SSLMode string

const (
	SSLModeDisable SSLMode = "disable"
	SSLModeRequire SSLMode = "require"
)

func ConnectAndMigrate(host, port, databaseName, user, password string, sslMode SSLMode) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, databaseName, sslMode)
	DB, err := sqlx.Open("postgres", connStr)

	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}
	NoteDB = DB
	return migrateUp(DB)
}

func migrateUp(db *sqlx.DB) error {
	db.Driver()
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrator.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", "./database/migrations"),
		os.Getenv("DB_NAME"), driver)

	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrator.ErrNoChange) {
		return err
	}
	return nil
}

func GetSSLMode() SSLMode {
	if SSLMode(os.Getenv("SSL_MODE")) == SSLModeRequire {
		return SSLModeRequire
	}
	return SSLModeDisable
}
