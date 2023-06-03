package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/jserva90/Erply-test-task/models"
)

func (m *SqliteDB) AddOrUpdateCustomerToDB(customerID int, clientCode, fullName, email, phone string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	checkStmt := `SELECT EXISTS(SELECT 1 FROM customers WHERE customer_id = ? AND client_code = ? LIMIT 1)`
	var exists bool
	err := m.DB.QueryRowContext(ctx, checkStmt, customerID, clientCode).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	now := time.Now().Unix()
	var stmt string
	if exists {
		stmt = `UPDATE customers
			SET full_name = ?, email = ?, phone = ?, created_at = ?
			WHERE customer_id = ? AND client_code = ?`
	} else {
		stmt = `INSERT INTO customers (client_code, customer_id, full_name, email, phone, created_at)
			VALUES (?,?,?,?,?,?)`
	}

	var args []interface{}
	if exists {
		args = []interface{}{fullName, email, phone, now, customerID, clientCode}
	} else {
		args = []interface{}{clientCode, customerID, fullName, email, phone, now}
	}

	_, err = m.DB.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}
	return nil
}

func (m *SqliteDB) GetCustomerFromDB(customerID, clientCode string) (*models.CustomerRecord, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	query := `SELECT customer_id, full_name, email, phone FROM customers WHERE customer_id = ? AND client_code = ?`

	row := m.DB.QueryRowContext(ctx, query, customerID, clientCode)
	var customer models.CustomerRecord

	err := row.Scan(
		&customer.CustomerID,
		&customer.FullName,
		&customer.Email,
		&customer.Phone,
	)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (m *SqliteDB) GetCustomerAddedTimestampFromDB(customerID, clientCode string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	query := `SELECT created_at FROM customers WHERE customer_id = ? AND client_code = ?`

	row := m.DB.QueryRowContext(ctx, query, customerID, clientCode)
	var createdAt int64

	err := row.Scan(
		&createdAt,
	)
	if err != nil {
		return 0, err
	}

	return createdAt, nil
}
