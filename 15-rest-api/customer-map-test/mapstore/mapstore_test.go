package mapstore

import (
	"errors"
	"go-training/customer-map/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapStore_Create(t *testing.T) {
	ms := NewMapStore()

	customer1 := domain.Customer{
		ID:    "cust101",
		Name:  "Krishnakumar",
		Email: "krishnakumarkp@gmail.com",
	}

	t.Run("Create a new customer", func(t *testing.T) {
		err := ms.Create(customer1)
		assert.NoError(t, err, "Expected no error")
	})

	t.Run("Create a duplicate customer", func(t *testing.T) {

		// Try to create a duplicate customer
		err := ms.Create(customer1)
		expectedErr := errors.New("Customer already presents")
		assert.EqualError(t, err, expectedErr.Error(), "Expected error 'Customer already present'")

	})

}

func TestMapStore_GetById(t *testing.T) {
	ms := NewMapStore()

	// Create a customer
	ms.Create(domain.Customer{ID: "1", Name: "John", Email: "john@test.com"})

	// Get the customer by ID
	c, err := ms.GetById("1")
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, "1", c.ID, "Expected customer ID to be '1'")
	assert.Equal(t, "John", c.Name, "Expected customer name to be 'John'")

	// Try to get a non-existent customer
	_, err = ms.GetById("2")
	expectedErr := errors.New("customer not present")
	assert.EqualError(t, err, expectedErr.Error(), "Expected error 'customer not present'")
}
