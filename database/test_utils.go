package database

import "database/sql"

func setupTestDB(dbPath string) (*SqliteDB, error) {
	var err error

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &SqliteDB{DB: db}, nil
}

func createTestCustomerTable(db *SqliteDB) error {
	tableStmt := `
	CREATE TABLE IF NOT EXISTS customers(
		client_code TEXT NOT NULL,
		customer_id INTEGER NOT NULL,
		full_name      TEXT,
		email TEXT,
		phone TEXT,
		created_at INTEGER
	)`

	_, err := db.DB.Exec(tableStmt)
	if err != nil {
		return err
	}

	return nil
}
