// with the standard testing package with basic conditional checks
// The convention in Go is that all test functions must:

// Be exported (i.e., start with an uppercase letter).
// Begin with the prefix Test.
// Accept a single parameter of type *testing.T.
// This means TestMapStore_Create and TestCreate are both valid names, as long as they start with Test.

// However, to improve clarity and consistency, especially in larger projects, it’s common to include the struct or function being tested as part of the name. For example:

// TestMapStore_Create (includes the struct being tested)
// TestCreate (shorter, but might be less descriptive in larger test suites)
// Example of Naming Conventions
// If you're testing a MapStore struct:

// Prefer TestMapStore_<MethodName> for clarity in identifying what is being tested.
// If the method belongs to a specific interface or struct, explicitly naming it can help:

// TestCustomerStore_Update (to indicate the interface)
package mapstore

import (
	"testing"

	"github.com/krishnakumarkp/customer-map/domain"
)

func TestMapStore_Create(t *testing.T) {
	store := NewMapStore()

	// Test creating a new customer
	c1 := domain.Customer{ID: "1", Name: "Alice", Email: "alice@example.com"}
	err := store.Create(c1)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Verify that the customer data was actually created in the mapstore
	createdCustomer, _ := store.GetById("1")
	if createdCustomer != c1 {
		t.Errorf("expected customer to be %v, but got %v", c1, createdCustomer)
	}

	// Test creating a customer with an existing ID
	err = store.Create(c1)
	if err == nil {
		t.Errorf("expected an error for duplicate ID, got nil")
	}
}

func TestMapStore_Update(t *testing.T) {
	store := NewMapStore()

	// Create an initial customer
	c1 := domain.Customer{ID: "1", Name: "Alice", Email: "alice@example.com"}
	store.Create(c1)

	// Test updating an existing customer
	cUpdated := domain.Customer{ID: "1", Name: "Alice Updated", Email: "alice.updated@example.com"}
	err := store.Update("1", cUpdated)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Verify that the customer data was actually updated
	updatedCustomer, _ := store.GetById("1")
	if updatedCustomer != cUpdated {
		t.Errorf("expected customer to be updated to %v, but got %v", cUpdated, updatedCustomer)
	}

	// Test updating a non-existing customer
	err = store.Update("2", cUpdated)
	if err == nil {
		t.Errorf("expected an error for non-existing customer, got nil")
	}
}

func TestMapStore_Delete(t *testing.T) {
	store := NewMapStore()

	// Create a customer to be deleted
	c1 := domain.Customer{ID: "1", Name: "Alice", Email: "alice@example.com"}
	store.Create(c1)

	// Test deleting an existing customer
	err := store.Delete("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Verify that the customer has been deleted
	_, err = store.GetById("1")
	if err == nil {
		t.Errorf("expected an error when retrieving a deleted customer, got nil")
	}

	// Test deleting a non-existing customer
	err = store.Delete("2")
	if err == nil {
		t.Errorf("expected an error for non-existing customer, got nil")
	}
}

func TestMapStore_GetById(t *testing.T) {
	store := NewMapStore()

	c1 := domain.Customer{ID: "1", Name: "Alice", Email: "alice@example.com"}
	store.Create(c1)

	// Test getting an existing customer by ID
	c, err := store.GetById("1")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if c != c1 {
		t.Errorf("expected customer %v, got %v", c1, c)
	}

	// Test getting a non-existing customer by ID
	_, err = store.GetById("2")
	if err == nil {
		t.Errorf("expected an error for non-existing customer, got nil")
	}
}
func TestMapStore_GetAll(t *testing.T) {
	store := NewMapStore()

	c1 := domain.Customer{ID: "1", Name: "Alice", Email: "alice@example.com"}
	c2 := domain.Customer{ID: "2", Name: "Bob", Email: "bob@example.com"}
	store.Create(c1)
	store.Create(c2)

	// Test getting all customers
	customers, err := store.GetAll()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(customers) != 2 {
		t.Errorf("expected 2 customers, got %d", len(customers))
	}
	if customers[0] != c1 && customers[1] != c1 {
		t.Errorf("expected customer %v in the list, but not found", c1)
	}
	if customers[0] != c2 && customers[1] != c2 {
		t.Errorf("expected customer %v in the list, but not found", c2)
	}
}
