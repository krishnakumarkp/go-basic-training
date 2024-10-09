package main

import (
	"errors"
	"testing"
)

// TestCustomerController tests the CustomerController methods
func TestCustomerController(t *testing.T) {
	// Create a new MapStore for testing
	store := NewMapStore()
	controller := CustomerController{
		Store: store,
	}

	// Define test customers
	customer1 := Customer{
		ID:    "cust101",
		Name:  "Krishnakumar",
		Email: "krishnakumarkp@gmail.com",
	}

	customer2 := Customer{
		ID:    "cust102",
		Name:  "Parvathy",
		Email: "parvathy@gmail.com",
	}

	// Test adding a customer
	controller.Add(customer1)
	// Check if customer has been added
	if _, err := store.GetById(customer1.ID); err != nil {
		t.Errorf("Failed to add customer: %v", err)
	}

	// Test updating a customer
	newName := "Krishna"
	controller.Update(Customer{
		ID:    customer1.ID,
		Name:  newName,
		Email: customer1.Email,
	})
	// Check if customer has been updated
	if c, _ := store.GetById(customer1.ID); c.Name != newName {
		t.Errorf("Failed to update customer name")
	}

	// Test deleting a customer
	controller.Delete(customer1.ID)
	// Check if customer has been deleted
	if _, err := store.GetById(customer1.ID); !errors.Is(err, errors.New("customer not present")) {
		t.Errorf("Failed to delete customer: %v", err)
	}

	// Test getting a customer
	store.Create(customer2)
	controller.Get(customer2.ID)
	// Check if customer is fetched
	if _, err := store.GetById(customer2.ID); err != nil {
		t.Errorf("Failed to get customer: %v", err)
	}

	// Test getting all customers
	customers, err := store.GetAll()
	if err != nil {
		t.Errorf("Failed to get all customers: %v", err)
	}
	if len(customers) != 1 {
		t.Errorf("Expected 1 customer, got %d", len(customers))
	}

	// Test deleting a non-existent customer
	controller.Delete("nonexistent")
	// Check if error is returned for non-existent customer
	if _, err := store.GetById("nonexistent"); !errors.Is(err, errors.New("customer not present")) {
		t.Errorf("Failed to handle non-existent customer deletion: %v", err)
	}

	// Test getting a non-existent customer
	controller.Get("nonexistent")
	// Check if error is returned for non-existent customer
	if _, err := store.GetById("nonexistent"); !errors.Is(err, errors.New("customer not present")) {
		t.Errorf("Failed to handle non-existent customer retrieval: %v", err)
	}
}
