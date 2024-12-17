package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/krishnakumarkp/customer-web/domain"
	"github.com/krishnakumarkp/customer-web/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAddCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockCustomerStore(ctrl)
	cc := CustomerController{Store: mockStore}

	// Test case 1: Valid customer data (Happy path)
	t.Run("Valid Customer", func(t *testing.T) {
		customer := domain.Customer{
			ID:    "1",
			Name:  "John Doe",
			Email: "john@test.com",
		}
		mockStore.EXPECT().Create(customer).Return(nil)

		body, _ := json.Marshal(customer)
		req := httptest.NewRequest("POST", "/customers", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		cc.Add(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.JSONEq(t, string(body), rec.Body.String())
	})
	// Test case 2: Invalid JSON body
	t.Run("Invalid JSON Body", func(t *testing.T) {
		invalidBody := `{"id": "1", "name": "John Doe", "email":}` // Invalid JSON

		req := httptest.NewRequest("POST", "/customers", bytes.NewBufferString(invalidBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		cc.Add(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "invalid character")
	})
	t.Run("Store Create Error", func(t *testing.T) {
		customer := domain.Customer{
			ID:    "1",
			Name:  "John Doe",
			Email: "john@test.com",
		}
		mockStore.EXPECT().Create(customer).Return(errors.New("database error"))

		body, _ := json.Marshal(customer)
		req := httptest.NewRequest("POST", "/customers", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		cc.Add(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "database error")
	})
}

func TestGetCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockCustomerStore(ctrl)
	cc := CustomerController{Store: mockStore}

	// Test case 1: Customer found (Happy path)
	t.Run("Customer Found", func(t *testing.T) {
		customer := domain.Customer{
			ID:    "1",
			Name:  "John Doe",
			Email: "john@test.com",
		}

		mockStore.EXPECT().GetById("1").Return(customer, nil)

		req := httptest.NewRequest("GET", "/customers/1", nil)
		req.SetPathValue("id", "1")

		rec := httptest.NewRecorder()

		cc.Get(rec, req)

		expectedBody, _ := json.Marshal(customer)
		assert.Equal(t, http.StatusFound, rec.Code)
		assert.JSONEq(t, string(expectedBody), rec.Body.String())
	})
	// Test case 2: Customer not found
	t.Run("Customer Not Found", func(t *testing.T) {
		mockStore.EXPECT().GetById("2").Return(domain.Customer{}, errors.New("customer not found"))

		req := httptest.NewRequest("GET", "/customers/2", nil)
		rec := httptest.NewRecorder()

		req.SetPathValue("id", "2")

		cc.Get(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "customer not found")
	})

	// Test case 4: Store returns an internal error
	t.Run("Store Internal Error", func(t *testing.T) {
		mockStore.EXPECT().GetById("3").Return(domain.Customer{}, errors.New("internal error"))

		req := httptest.NewRequest("GET", "/customers/3", nil)
		rec := httptest.NewRecorder()

		// Manually extract path variable and inject it into the request
		req.SetPathValue("id", "3")

		cc.Get(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "internal error")
	})

}

func TestGetAllCustomers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockCustomerStore(ctrl)
	controller := CustomerController{Store: mockStore}

	// Test case 1: Successfully retrieve all customers
	t.Run("GetAll Customers Successfully", func(t *testing.T) {
		customers := []domain.Customer{
			{ID: "1", Name: "John Doe", Email: "john@test.com"},
			{ID: "2", Name: "Jane Smith", Email: "jane@test.com"},
		}

		mockStore.EXPECT().GetAll().Return(customers, nil)

		req := httptest.NewRequest("GET", "/customers", nil)
		rr := httptest.NewRecorder()

		http.HandlerFunc(controller.GetAll).ServeHTTP(rr, req)

		expectedBody, _ := json.Marshal(customers)

		assert.Equal(t, http.StatusFound, rr.Code)
		assert.JSONEq(t, string(expectedBody), rr.Body.String())
	})

	// Test case 2: Error retrieving customers
	t.Run("GetAll Customers Error", func(t *testing.T) {
		mockStore.EXPECT().GetAll().Return(nil, errors.New("database error"))

		req := httptest.NewRequest("GET", "/customers", nil)
		rr := httptest.NewRecorder()

		http.HandlerFunc(controller.GetAll).ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Contains(t, rr.Body.String(), "database error")
	})

	// Test case 3: No customers found (empty list)
	t.Run("GetAll No Customers", func(t *testing.T) {
		mockStore.EXPECT().GetAll().Return([]domain.Customer{}, nil)

		req := httptest.NewRequest("GET", "/customers", nil)
		rr := httptest.NewRecorder()

		http.HandlerFunc(controller.GetAll).ServeHTTP(rr, req)

		expectedBody := "[]"

		assert.Equal(t, http.StatusFound, rr.Code)
		assert.JSONEq(t, expectedBody, rr.Body.String())
	})
}
