### **Introduction to sqlmock**

`sqlmock` is a powerful library for mocking database operations in unit tests. It allows you to simulate SQL queries and interactions without needing an actual database, making it faster and more controlled than running tests against a real database.

In this tutorial, we will walk through how to use `sqlmock` to test the CRUD operations of a `CustomerStore` in a Go application.

### **Setting Up the Project**

Let's assume you have a Go application that interacts with a MySQL database and has a `CustomerStore` that performs CRUD operations on a `customer` table. Your goal is to unit test these methods by mocking the database interactions.

1. **Install `sqlmock`**:
```
go get github.com/DATA-DOG/go-sqlmock
```


Test Case 1: Testing the `Create` Method

The `Create` method inserts a new customer into the database.
To test this, we mock the `INSERT` statement and verify that the query is executed as expected.

```
func TestCustomerStore_Create(t *testing.T) {
	// Set up the mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error occurred while creating mock database: %v", err)
	}
	defer db.Close()

	// Create a DataStore instance
	dataStore := DataStore{Db: db}
	customerStore := NewCustomerStore(dataStore)

	// Define the customer to be created
	customer := domain.Customer{ID: "1", Name: "Alice", Email: "alice@example.com"}

	// Mock the expected behavior: prepare the query and execute it
	mock.ExpectPrepare("INSERT INTO customer").ExpectExec().
		WithArgs(customer.ID, customer.Name, customer.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Test the Create method
	err = customerStore.Create(customer)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Ensure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}
```
In this test:

- We create a mock database with `sqlmock.New()`.
- We mock the `Prepare` method to expect an `INSERT INTO customer` query.
- We use `WillReturnResult` to simulate a successful insertion.
- After calling the `Create` method, we check for errors and verify that the expectations are met with `mock.ExpectationsWereMet()`.


for other testas look at customer_store_test.go