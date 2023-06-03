package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jserva90/Erply-test-task/models"
	_ "github.com/mattes/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AddOrUpdateCustomerToDB(customerID int, clientCode, fullName, email, phone string) error
	GetCustomerFromDB(customerID, clientCode string) (*models.CustomerRecord, error)
	AddSession(clientCode, username, password, sessionKey, sessionToken string) error
	GetSession(sessionToken string) (*models.Session, error)
	RemoveSession(sessionToken string) error
	GetCustomerAddedTimestampFromDB(customerID, clientCode string) (int64, error)
}

type SqliteDB struct {
	DB *sql.DB
}

const DbTimeout = time.Second * 3

func (m *SqliteDB) Connection() *sql.DB {
	return m.DB
}

func OpenDB(dbPath string) (*sql.DB, error) {
	var err error

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = migrateDB(db, "database/migrations")
	if err != nil {
		log.Fatalf("Error applying database migrations: %v", err)
	}

	log.Println("Database connection established")

	return db, nil
}

func migrateDB(db *sql.DB, migrationsDir string) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("failed to initialize sqlite3 driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsDir,
		"sqlite3", driver)
	if err != nil {
		return fmt.Errorf("failed to initialize migrate instance: %w", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply database migrations: %w", err)
	}

	return nil
}
