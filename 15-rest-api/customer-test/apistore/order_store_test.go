package apistore

import (
	"encoding/json"
	"errors"
	"go-training/customer-map/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrdersByCustomerID(t *testing.T) {

	t.Run("Get orders for customer123", func(t *testing.T) {
		// Create a mock server
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check the URL called by the client
			if r.URL.String() != "/customer/customer123/order" {
				t.Errorf("unexpected URL called: %s", r.URL.String())
			}

			// Mock response
			orders := []domain.Order{{BookID: "1234", Quantity: 1}}
			json.NewEncoder(w).Encode(orders)
		}))
		defer mockServer.Close()

		// Create an instance of OrderStore with the mock server's URL
		orderStore := OrderStore{BaseURL: mockServer.URL}

		// Call the method being tested
		orders, err := orderStore.GetOrdersByCustomerID("customer123")

		// Assert that there is no error
		assert.NoError(t, err, "unexpected error")

		// Assert that the expected order is returned
		expectedOrder := domain.Order{BookID: "1234", Quantity: 1}
		assert.Len(t, orders, 1, "unexpected number of orders returned")
		assert.Equal(t, expectedOrder, orders[0], "unexpected orders returned")
	})

	t.Run("Get orders for non-existing customer", func(t *testing.T) {
		// Create a mock server
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check the URL called by the client
			if r.URL.String() != "/customer/customer124/order" {
				t.Errorf("unexpected URL called: %s", r.URL.String())
			}

			w.WriteHeader(http.StatusNotFound)
		}))
		defer mockServer.Close()

		// Create an instance of OrderStore with the mock server's URL
		orderStore := OrderStore{BaseURL: mockServer.URL}
		// Call the method being tested for a non-existing customer
		_, err := orderStore.GetOrdersByCustomerID("customer124")

		expectedError := errors.New("failed to get orders")
		// Assert that there is an error
		assert.Error(t, err, expectedError.Error())
	})

}
