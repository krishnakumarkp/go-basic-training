package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCustomerController struct {
	mock.Mock
}

func (cc MockCustomerController) Add(w http.ResponseWriter, r *http.Request) {
	cc.Called(w, r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"id": "cust105", "name": "arun", "email": "arun@gmail.com"}`))
}

func (cc MockCustomerController) Get(w http.ResponseWriter, r *http.Request) {
	cc.Called(w, r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write([]byte(`{"id": "cust105", "name": "arun", "email": "arun@gmail.com"}`))

}

func (cc MockCustomerController) Update(w http.ResponseWriter, r *http.Request) {
	cc.Called(w, r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write([]byte(`{"id": "cust105", "name": "arun", "email": "arun@gmail.com"}`))

}

func (cc MockCustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	cc.Called(w, r)
	w.WriteHeader(http.StatusNoContent)
}

func (cc MockCustomerController) GetAll(w http.ResponseWriter, r *http.Request) {
	cc.Called(w, r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write([]byte(`[{"id": "cust105", "name": "arun", "email": "arun@gmail.com"}]`))
}

func TestGetRouterWithMockedController(t *testing.T) {
	// Create a mock controller
	mockController := new(MockCustomerController)

	// Set up the router
	router := GetRouter(mockController)

	// Define test cases
	testCases := []struct {
		name               string
		method             string
		path               string
		body               string
		mockSetup          func()
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:   "Valid POST /customer",
			method: "POST",
			path:   "/customer",
			body:   `{"id": "cust105", "name": "arun", "email": "arun@gmail.com"}`,
			mockSetup: func() {
				//mockController.On("Add", mock.AnythingOfType("http.ResponseWriter"), mock.AnythingOfType("*http.Request")).Return()
				mockController.On("Add", mock.Anything, mock.AnythingOfType("*http.Request")).Return()
			},
			expectedStatusCode: http.StatusCreated,
			expectedResponse:   `{"id": "cust105", "name": "arun", "email": "arun@gmail.com"}`,
		},
		{
			name:   "Valid GET /customer/{id}",
			method: "GET",
			path:   "/customer/cust105",
			body:   "",
			mockSetup: func() {
				mockController.On("Get", mock.Anything, mock.Anything).Return()
			},
			expectedStatusCode: http.StatusFound,
			expectedResponse:   `{"id": "cust105", "name": "arun", "email": "arun@gmail.com"}`,
		},
		{
			name:   "Valid PUT /customer/{id}",
			method: "PUT",
			path:   "/customer/cust105",
			body:   `{"id": "cust105", "name": "arun", "email": "arun@gmail.com"}`,
			mockSetup: func() {
				mockController.On("Update", mock.Anything, mock.Anything).Return()
			},
			expectedStatusCode: http.StatusFound,
			expectedResponse:   `{"id": "cust105", "name": "arun", "email": "arun@gmail.com"}`,
		},
		{
			name:   "Valid DELETE /customer/{id}",
			method: "DELETE",
			path:   "/customer/cust105",
			body:   "",
			mockSetup: func() {
				mockController.On("Delete", mock.Anything, mock.Anything).Return()
			},
			expectedStatusCode: http.StatusNoContent,
			expectedResponse:   "",
		},
		{
			name:   "Valid GET /customer",
			method: "GET",
			path:   "/customer",
			body:   "",
			mockSetup: func() {
				mockController.On("GetAll", mock.Anything, mock.Anything).Return()
			},
			expectedStatusCode: http.StatusFound,
			expectedResponse:   `[{"id": "cust105", "name": "arun", "email": "arun@gmail.com"}]`,
		},
	}

	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set up mock behavior
			if tc.mockSetup != nil {
				tc.mockSetup()
			}

			// Create the request
			req := httptest.NewRequest(tc.method, tc.path, bytes.NewBufferString(tc.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Serve the request using the router
			router.ServeHTTP(w, req)

			// Verify the response
			assert.Equal(t, tc.expectedStatusCode, w.Code)
			if tc.expectedResponse != "" {
				assert.Contains(t, w.Body.String(), tc.expectedResponse)
			}

			// Assert mock expectations
			mockController.AssertExpectations(t)
		})
	}
}
