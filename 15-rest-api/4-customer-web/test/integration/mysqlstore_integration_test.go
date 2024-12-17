package mysqlstore

import (
	"testing"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/krishnakumarkp/customer-web/apputil"
	"github.com/krishnakumarkp/customer-web/domain"
	"github.com/krishnakumarkp/customer-web/mysqlstore"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() mysqlstore.DataStore {
	config := mysqlstore.Config{
		Host:     apputil.AppConfig.DBHost,
		Port:     apputil.AppConfig.DBPort,
		User:     apputil.AppConfig.DBUser,
		Password: apputil.AppConfig.DBPassword,
		Database: apputil.AppConfig.Database,
	}
	// Creates a Mysql DB instance
	dbStore, err := mysqlstore.New(config)
	if err != nil {
		panic(err)
	}
	return dbStore
}

func TestCustomerStore_Create(t *testing.T) {
	// Set up test database
	db := setupTestDB()

	// Initialize the store with the real DB connection
	store := mysqlstore.NewCustomerStore(db)

	// Prepare the customer data for the test
	customer := domain.Customer{
		ID:    "cust123",
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	// Test Create method
	err := store.Create(customer)

	// Verify that there was no error
	assert.Nil(t, err)

	// Verify the customer was actually added to the database
	var fetchedCustomer domain.Customer
	err = db.QueryRow("SELECT id, name, email FROM customer WHERE id = ?", customer.ID).
		Scan(&fetchedCustomer.ID, &fetchedCustomer.Name, &fetchedCustomer.Email)
	assert.Nil(t, err)
	assert.Equal(t, customer.ID, fetchedCustomer.ID)
	assert.Equal(t, customer.Name, fetchedCustomer.Name)
	assert.Equal(t, customer.Email, fetchedCustomer.Email)
}

func TestCustomerStore_GetById(t *testing.T) {
	// Set up test database
	db := setupTestDB()

	// Initialize the store with the real DB connection
	store := NewCustomerStore(db)

	// Prepare a customer to insert into the database
	customer := domain.Customer{
		ID:    "cust124",
		Name:  "Jane Doe",
		Email: "janedoe@example.com",
	}

	// Insert the customer into the database
	err := store.Create(customer)
	assert.Nil(t, err)

	// Test GetById method
	fetchedCustomer, err := store.GetById(customer.ID)
	assert.Nil(t, err)
	assert.Equal(t, customer.ID, fetchedCustomer.ID)
	assert.Equal(t, customer.Name, fetchedCustomer.Name)
	assert.Equal(t, customer.Email, fetchedCustomer.Email)
}

func TestCustomerStore_GetAll(t *testing.T) {
	// Set up test database
	db := setupTestDB()
	defer db.Close()

	// Initialize the store with the real DB connection
	store := NewCustomerStore(db)

	// Prepare customers to insert into the database
	customers := []domain.Customer{
		{ID: "cust125", Name: "Alice", Email: "alice@example.com"},
		{ID: "cust126", Name: "Bob", Email: "bob@example.com"},
	}

	// Insert the customers into the database
	for _, customer := range customers {
		err := store.Create(customer)
		assert.Nil(t, err)
	}

	// Test GetAll method
	allCustomers, err := store.GetAll()
	assert.Nil(t, err)
	assert.Len(t, allCustomers, 2)

	// Check that the customers are correctly fetched
	assert.Contains(t, allCustomers, customers[0])
	assert.Contains(t, allCustomers, customers[1])
}

func TestCustomerStore_Update(t *testing.T) {
	// Set up test database
	db := setupTestDB()
	defer db.Close()

	// Initialize the store with the real DB connection
	store := NewCustomerStore(db)

	// Prepare a customer to insert into the database
	customer := domain.Customer{
		ID:    "cust127",
		Name:  "Charlie",
		Email: "charlie@example.com",
	}

	// Insert the customer into the database
	err := store.Create(customer)
	assert.Nil(t, err)

	// Prepare updated customer data
	updatedCustomer := domain.Customer{
		ID:    customer.ID,
		Name:  "Charlie Updated",
		Email: "charlie_updated@example.com",
	}

	// Test Update method
	err = store.Update(customer.ID, updatedCustomer)
	assert.Nil(t, err)

	// Verify that the customer was updated
	fetchedCustomer, err := store.GetById(customer.ID)
	assert.Nil(t, err)
	assert.Equal(t, updatedCustomer.Name, fetchedCustomer.Name)
	assert.Equal(t, updatedCustomer.Email, fetchedCustomer.Email)
}

func TestCustomerStore_Delete(t *testing.T) {
	// Set up test database
	db := setupTestDB()
	defer db.Close()

	// Initialize the store with the real DB connection
	store := NewCustomerStore(db)

	// Prepare a customer to insert into the database
	customer := domain.Customer{
		ID:    "cust128",
		Name:  "David",
		Email: "david@example.com",
	}

	// Insert the customer into the database
	err := store.Create(customer)
	assert.Nil(t, err)

	// Test Delete method
	err = store.Delete(customer.ID)
	assert.Nil(t, err)

	// Verify that the customer was deleted
	_, err = store.GetById(customer.ID)
	assert.NotNil(t, err)
}
