package database

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jserva90/Erply-test-task/models"
	_ "github.com/mattn/go-sqlite3"
)

func TestAddOrUpdateCustomerToDB(t *testing.T) {
	db, err := setupTestDB("./database_test.db")
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
	}

	defer func() {
		db.DB.Close()
		os.Remove("./database_test.db")
	}()

	err = createTestCustomerTable(db)
	if err != nil {
		t.Errorf("Failed to create table: %s", err)
	}

	customerID := 1
	clientCode := "ABC"
	fullName := "John Doe"
	email := "john@example.com"
	phone := "1234567890"

	err = db.AddOrUpdateCustomerToDB(customerID, clientCode, fullName, email, phone)
	if err != nil {
		t.Errorf("Failed to add customer to the database: %s", err)
	}

	fullName = "John Smith"
	email = "john.smith@example.com"
	phone = "9876543210"

	err = db.AddOrUpdateCustomerToDB(customerID, clientCode, fullName, email, phone)
	if err != nil {
		t.Errorf("Failed to update customer in the database: %s", err)
	}
}

func TestGetCustomerFromDB(t *testing.T) {
	db, err := setupTestDB("./database_test.db")
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
	}

	defer func() {
		db.DB.Close()
		os.Remove("./database_test.db")
	}()

	err = createTestCustomerTable(db)
	if err != nil {
		t.Errorf("Failed to create table: %s", err)
	}

	customerID := 1
	clientCode := "ABC"
	fullName := "John Doe"
	email := "john@example.com"
	phone := "1234567890"
	err = db.AddOrUpdateCustomerToDB(customerID, clientCode, fullName, email, phone)
	if err != nil {
		t.Fatalf("Failed to insert test customer: %s", err)
	}

	customer, err := db.GetCustomerFromDB(fmt.Sprint(customerID), clientCode)
	if err != nil {
		t.Errorf("Failed to retrieve customer from the database: %s", err)
	}

	expectedCustomer := &models.CustomerRecord{
		CustomerID: customerID,
		FullName:   fullName,
		Email:      email,
		Phone:      phone,
	}
	if customer.CustomerID != expectedCustomer.CustomerID ||
		customer.FullName != expectedCustomer.FullName ||
		customer.Email != expectedCustomer.Email ||
		customer.Phone != expectedCustomer.Phone {
		t.Errorf("Retrieved customer does not match the expected customer")
	}
}

func TestGetCustomerAddedTimestampFromDB(t *testing.T) {
	db, err := setupTestDB("./database_test.db")
	if err != nil {
		t.Errorf("failed to setup test database: %v", err)
	}

	defer func() {
		db.DB.Close()
		os.Remove("./database_test.db")
	}()

	err = createTestCustomerTable(db)
	if err != nil {
		t.Errorf("Failed to create table: %s", err)
	}

	customerID := 1
	clientCode := "ABC"
	fullName := "John Doe"
	email := "john@example.com"
	phone := "1234567890"
	err = db.AddOrUpdateCustomerToDB(customerID, clientCode, fullName, email, phone)
	if err != nil {
		t.Fatalf("Failed to insert test customer: %s", err)
	}

	createdAt, err := db.GetCustomerAddedTimestampFromDB(fmt.Sprint(customerID), clientCode)
	if err != nil {
		t.Errorf("Failed to retrieve customer's added timestamp from the database: %s", err)
	}

	currentTime := time.Now().Unix()
	if currentTime-createdAt > 2 {
		t.Errorf("Retrieved customer's added timestamp is not within the expected range")
	}
}
