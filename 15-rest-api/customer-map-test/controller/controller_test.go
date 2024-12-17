package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-training/customer-map/domain"
	"go-training/customer-map/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCustomerController(t *testing.T) {

	t.Run("Add customer", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockStore := mocks.NewMockCustomerStore(ctrl)

		customer1 := domain.Customer{
			ID:    "cust101",
			Name:  "Krishnakumar",
			Email: "krishnakumarkp@gmail.com",
		}

		mockStore.EXPECT().Create(customer1).Return(nil)

		controller := NewCustomerController(mockStore)

		// Marshal the sample customer object to JSON
		customerJSON, err := json.Marshal(customer1)
		if err != nil {
			t.Fatalf("Failed to marshal customer to JSON: %v", err)
		}
		// Create a new HTTP request with POST method and customer JSON data
		req, err := http.NewRequest("POST", "/customer", bytes.NewBuffer(customerJSON))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()

		// Call the Add method with the request and recorder
		http.HandlerFunc(controller.Add).ServeHTTP(rr, req)

		// Check the status code
		assert.Equal(t, http.StatusCreated, rr.Code, "handler returned wrong status code")

		// Check the response body
		expected := string(customerJSON)
		assert.Equal(t, expected, rr.Body.String(), "handler returned unexpected body")

	})

	t.Run("Add duplicate customer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockStore := mocks.NewMockCustomerStore(ctrl)

		customer1 := domain.Customer{
			ID:    "cust101",
			Name:  "Krishnakumar",
			Email: "krishnakumarkp@gmail.com",
		}

		mockStore.EXPECT().Create(customer1).Return(errors.New("Error 1062 (23000): Duplicate entry 'cust101' for key 'customer.ID'"))

		controller := NewCustomerController(mockStore)

		// Marshal the sample customer object to JSON
		customerJSON, err := json.Marshal(customer1)
		if err != nil {
			t.Fatalf("Failed to marshal customer to JSON: %v", err)
		}
		// Create a new HTTP request with POST method and customer JSON data
		req, err := http.NewRequest("POST", "/customer", bytes.NewBuffer(customerJSON))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()

		// Call the Add method with the request and recorder
		http.HandlerFunc(controller.Add).ServeHTTP(rr, req)

		// Check the status code
		assert.Equal(t, http.StatusInternalServerError, rr.Code, "handler returned wrong status code")

		// Check the response body
		expected := "Error 1062 (23000): Duplicate entry 'cust101' for key 'customer.ID'\n"
		assert.Equal(t, expected, rr.Body.String(), "handler returned unexpected body")
	})

}

func dieOn(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}
