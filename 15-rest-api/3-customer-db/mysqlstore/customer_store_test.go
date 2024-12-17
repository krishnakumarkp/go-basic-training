package mysqlstore

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/krishnakumarkp/customer-db/domain"
)

func TestCustomerStore_Create(t *testing.T) {
	// Set up the mock database and mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error occurred while creating mock database: %v", err)
	}
	defer db.Close()

	// Create a DataStore instance
	dataStore := DataStore{Db: db}
	customerStore := NewCustomerStore(dataStore)

	// Mock the expected behavior of the database interaction
	customer := domain.Customer{ID: "1", Name: "Alice", Email: "alice@example.com"}

	mock.ExpectPrepare("INSERT INTO customer").ExpectExec().
		WithArgs(customer.ID, customer.Name, customer.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Test Create method
	err = customerStore.Create(customer)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Ensure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}

func TestCustomerStore_GetById(t *testing.T) {
	// Set up the mock database and mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error occurred while creating mock database: %v", err)
	}
	defer db.Close()

	// Create a DataStore instance
	dataStore := DataStore{Db: db}
	customerStore := NewCustomerStore(dataStore)

	// Successful case
	t.Run("Success", func(t *testing.T) {
		customer := domain.Customer{ID: "1", Name: "Alice", Email: "alice@example.com"}

		mock.ExpectQuery("select id, name, email FROM customer where ID = \\?").
			WithArgs(customer.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).
				AddRow(customer.ID, customer.Name, customer.Email))

		result, err := customerStore.GetById(customer.ID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if result != customer {
			t.Errorf("expected %v, got %v", customer, result)
		}
	})

	// Error case: No rows found
	t.Run("No rows found", func(t *testing.T) {
		mock.ExpectQuery("select id, name, email FROM customer where ID = \\?").
			WithArgs("2").
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}))

		_, err := customerStore.GetById("2")
		if err == nil {
			t.Errorf("expected error, got none")
		} else if err.Error() != "customer not found" {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Error case: Query error
	t.Run("Query error", func(t *testing.T) {
		mock.ExpectQuery("select id, name, email FROM customer where ID = \\?").
			WithArgs("3").
			WillReturnError(fmt.Errorf("query error"))

		_, err := customerStore.GetById("3")
		if err == nil {
			t.Errorf("expected error, got none")
		} else if err.Error() != "query error" {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func TestCustomerStore_GetAll(t *testing.T) {
	// Set up the mock database and mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error occurred while creating mock database: %v", err)
	}
	defer db.Close()

	// Create a DataStore instance
	dataStore := DataStore{Db: db}
	customerStore := NewCustomerStore(dataStore)

	// Successful case: Retrieve multiple customers
	t.Run("Success", func(t *testing.T) {
		customers := []domain.Customer{
			{ID: "1", Name: "Alice", Email: "alice@example.com"},
			{ID: "2", Name: "Bob", Email: "bob@example.com"},
		}

		rows := sqlmock.NewRows([]string{"id", "name", "email"}).
			AddRow(customers[0].ID, customers[0].Name, customers[0].Email).
			AddRow(customers[1].ID, customers[1].Name, customers[1].Email)

		mock.ExpectQuery("select id, name, email FROM customer ORDER BY id;").
			WillReturnRows(rows)

		result, err := customerStore.GetAll()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if len(result) != len(customers) {
			t.Errorf("expected %d customers, got %d", len(customers), len(result))
		}
		for i, customer := range customers {
			if result[i] != customer {
				t.Errorf("expected %v, got %v", customer, result[i])
			}
		}
	})

	// Error case: Query fails
	t.Run("Query error", func(t *testing.T) {
		mock.ExpectQuery("select id, name, email FROM customer ORDER BY id;").
			WillReturnError(fmt.Errorf("query error"))

		_, err := customerStore.GetAll()
		if err == nil {
			t.Errorf("expected error, got none")
		} else if err.Error() != "could not retrieve customers" {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// // Error case: Scan error
	// t.Run("Scan error", func(t *testing.T) {
	// 	rows := sqlmock.NewRows([]string{"id", "name", "email"}).
	// 		AddRow("1", "Alice", "alice@example.com").
	// 		AddRow("invalid", "Bob", "bob@example.com") // Invalid ID causing a scan error

	// 	mock.ExpectQuery("select id, name, email FROM customer ORDER BY id;").
	// 		WillReturnRows(rows)

	// 	_, err := customerStore.GetAll()
	// 	if err == nil {
	// 		t.Errorf("expected error, got none")
	// 	}
	// })

	// Empty case: No customers in the database
	t.Run("No customers", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "email"})

		mock.ExpectQuery("select id, name, email FROM customer ORDER BY id;").
			WillReturnRows(rows)

		result, err := customerStore.GetAll()
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if len(result) != 0 {
			t.Errorf("expected no customers, got %d", len(result))
		}
	})
}
