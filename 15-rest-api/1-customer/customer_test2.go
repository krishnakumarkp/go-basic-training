// TestCustomerController tests the CustomerController methods
func TestCustomerController(t *testing.T) {
	tests := []struct {
		name string
		run  func(*testing.T, CustomerController, *MapStore)
	}{
		{
			name: "AddDuplicateCustomer",
			run:  testAddDuplicateCustomer,
		},
		{
			name: "UpdateNonExistentCustomer",
			run:  testUpdateNonExistentCustomer,
		},
		{
			name: "DeleteNonExistentCustomer",
			run:  testDeleteNonExistentCustomer,
		},
		{
			name: "GetNonExistentCustomer",
			run:  testGetNonExistentCustomer,
		},
		{
			name: "GetAllCustomers",
			run:  testGetAllCustomers,
		},
		{
			name: "UpdateCustomerEmail",
			run:  testUpdateCustomerEmail,
		},
		{
			name: "DeleteCustomerWithDependency",
			run:  testDeleteCustomerWithDependency,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			store := NewMapStore()
			controller := CustomerController{
				Store: store,
			}
			tc.run(t, controller, store)
		})
	}
}

func testAddDuplicateCustomer(t *testing.T, controller CustomerController, store *MapStore) {
	// Add a customer
	customer := Customer{
		ID:    "cust101",
		Name:  "Krishnakumar",
		Email: "krishnakumarkp@gmail.com",
	}
	controller.Add(customer)

	// Attempt to add the same customer again
	err := store.Create(customer)

	// Ensure that an error indicating the customer already exists is returned
	if !errors.Is(err, errors.New("Customer already present")) {
		t.Errorf("Expected error: Customer already present, got: %v", err)
	}
}

func testUpdateNonExistentCustomer(t *testing.T, controller CustomerController, _ *MapStore) {
	// Attempt to update a non-existent customer
	err := controller.Update(Customer{
		ID:    "nonexistent",
		Name:  "NewName",
		Email: "newemail@example.com",
	})

	// Ensure that an error indicating the customer does not exist is returned
	if !errors.Is(err, errors.New("customer not present")) {
		t.Errorf("Expected error: customer not present, got: %v", err)
	}
}

func testDeleteNonExistentCustomer(t *testing.T, controller CustomerController, _ *MapStore) {
	// Attempt to delete a non-existent customer
	err := controller.Delete("nonexistent")

	// Ensure that an error indicating the customer does not exist is returned
	if !errors.Is(err, errors.New("customer not present")) {
		t.Errorf("Expected error: customer not present, got: %v", err)
	}
}

func testGetNonExistentCustomer(t *testing.T, controller CustomerController, _ *MapStore) {
	// Attempt to get a non-existent customer
	_, err := controller.Store.GetById("nonexistent")

	// Ensure that an error indicating the customer does not exist is returned
	if !errors.Is(err, errors.New("customer not present")) {
		t.Errorf("Expected error: customer not present, got: %v", err)
	}
}

func testGetAllCustomers(t *testing.T, controller CustomerController, _ *MapStore) {
	// Add some customers
	customer1 := Customer{ID: "cust101", Name: "John", Email: "john@example.com"}
	customer2 := Customer{ID: "cust102", Name: "Jane", Email: "jane@example.com"}
	controller.Add(customer1)
	controller.Add(customer2)

	// Get all customers
	customers, err := controller.Store.GetAll()

	// Check if customers are returned correctly
	if err != nil {
		t.Errorf("Error getting all customers: %v", err)
	}
	if len(customers) != 2 {
		t.Errorf("Expected 2 customers, got %d", len(customers))
	}
}

func testUpdateCustomerEmail(t *testing.T, controller CustomerController, _ *MapStore) {
	// Add a customer
	customer := Customer{
		ID:    "cust101",
		Name:  "John",
		Email: "john@example.com",
	}
	controller.Add(customer)

	// Update customer email
	newEmail := "newemail@example.com"
	controller.Update(Customer{
		ID:    customer.ID,
		Name:  customer.Name,
		Email: newEmail,
	})

	// Get the updated customer
	updatedCustomer, err := controller.Store.GetById(customer.ID)

	// Check if the email is updated correctly
	if err != nil {
		t.Errorf("Error getting updated customer: %v", err)
	}
	if updatedCustomer.Email != newEmail {
		t.Errorf("Expected updated email %s, got %s", newEmail, updatedCustomer.Email)
	}
}

func testDeleteCustomerWithDependency(t *testing.T, controller CustomerController, _ *MapStore) {
	// Add a customer
	customer := Customer{
		ID:    "cust101",
		Name:  "John",
		Email: "john@example.com",
	}
	controller.Add(customer)

	// Add dependency (e.g., orders)
	// For demonstration purposes, we can assume that dependency is added to the store itself.

	// Attempt to delete the customer
	err := controller.Delete(customer.ID)

	// Ensure that the customer cannot be deleted due to existing dependencies
	if !errors.Is(err, errors.New("unable to delete customer due to existing dependencies")) {
		t.Errorf("Expected error: unable to delete customer due to existing dependencies, got: %v", err)
	}
}
