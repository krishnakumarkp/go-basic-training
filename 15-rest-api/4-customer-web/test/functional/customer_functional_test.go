// test/functional/customer_functional_test.go
package functional

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/krishnakumarkp/customer-web/app"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Setenv("APP_ENV", "test")
	go app.StartApp()
	time.Sleep(5 * time.Second)
	os.Exit(m.Run())

}

func TestCustomerAPI(t *testing.T) {
	baseURL := "http://localhost:8080/customer"

	t.Run("Add Customer", func(t *testing.T) {
		customer := map[string]string{
			"id":    "1",
			"name":  "John Doe",
			"email": "john@test.com",
		}
		body, _ := json.Marshal(customer)

		resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(body))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("Retrieve Customer", func(t *testing.T) {
		resp, err := http.Get(baseURL + "/1")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var retrievedCustomer map[string]string
		json.NewDecoder(resp.Body).Decode(&retrievedCustomer)

		expectedCustomer := map[string]string{
			"id":    "1",
			"name":  "John Doe",
			"email": "john@test.com",
		}
		assert.Equal(t, expectedCustomer, retrievedCustomer)
	})

	t.Run("Update Customer", func(t *testing.T) {
		updatedCustomer := map[string]string{
			"id":    "1",
			"name":  "Jane Doe",
			"email": "jane@test.com",
		}
		body, _ := json.Marshal(updatedCustomer)

		req, err := http.NewRequest(http.MethodPut, baseURL+"/1", bytes.NewBuffer(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusFound, resp.StatusCode)

		// Verify the update
		resp, err = http.Get(baseURL + "/1")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var retrievedCustomer map[string]string
		json.NewDecoder(resp.Body).Decode(&retrievedCustomer)
		assert.Equal(t, updatedCustomer, retrievedCustomer)
	})

	t.Run("Delete Customer", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, baseURL+"/1", nil)
		assert.NoError(t, err)

		client := &http.Client{}
		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)

		// Verify the deletion
		resp, err = http.Get(baseURL + "/1")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
