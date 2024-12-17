package controller

import (
	"bytes"
	"encoding/json"
	"go-training/customer-map/domain"
	"net/http"
	"os"
	"testing"
	"time"

	"go-training/customer-map/app"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv("APP_ENV", "test")
	go app.StartApp()
	time.Sleep(5 * time.Second)
	os.Exit(m.Run())
}

func TestCustomerCreate(t *testing.T) {

	t.Run("Add customer", func(t *testing.T) {

		customer := domain.Customer{
			ID:    "cust1111",
			Name:  "Krishnakumar",
			Email: "krishnakumarkp@gmail.com",
		}

		// Marshal the expected customer data to JSON
		expectedJSON, err := json.Marshal(customer)
		assert.NoError(t, err, "Failed to marshal expected customer data to JSON")

		// Make a POST request to the local server
		resp, err := http.Post("http://localhost:8081/customer", "application/json", bytes.NewBuffer(expectedJSON))
		assert.NoError(t, err, "Failed to make POST request")
		defer resp.Body.Close()

		// Check the status code
		assert.Equal(t, http.StatusCreated, resp.StatusCode, "handler returned wrong status code")

		var responseCustomer domain.Customer

		err = json.NewDecoder(resp.Body).Decode(&responseCustomer)
		assert.NoError(t, err, "Failed to decode response body")

		assert.Equal(t, customer.ID, responseCustomer.ID, "ID mismatch")
		assert.Equal(t, customer.Name, responseCustomer.Name, "Name mismatch")
		assert.Equal(t, customer.Email, responseCustomer.Email, "Email mismatch")

	})

}
