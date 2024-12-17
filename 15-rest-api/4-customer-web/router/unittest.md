### Mocking Controllers with `github.com/stretchr/testify/mock`

In this tutorial, we will guide you through using `github.com/stretchr/testify/mock` to mock a controller in Go. This approach is ideal for unit testing, allowing you to isolate dependencies and test the logic of your code without actually invoking the real controller methods.

We will demonstrate mocking the `CustomerController` interface and testing its routing functionality.

### Step 1: Install Dependencies

To get started, you need to install the required dependencies, including `github.com/stretchr/testify` for mocking and assertions.

```
go get github.com/stretchr/testify

```

This will install the `testify` package, which contains helpful utilities like mock objects and assertions.

### Step 2: Define Your Controller Interface

In this tutorial, we assume you have a `CustomerControllerInterface` interface in your `controller` package, which provides methods for handling requests related to customer data. Here's an example:

```
package controller

import "net/http"

type CustomerControllerInterface interface {
	Add(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

```

This interface includes methods for handling basic HTTP operations: `Add`, `Get`, `Update`, `Delete`, and `GetAll`.

### Step 3: Define the Mock Controller

The `testify/mock` package provides the ability to easily mock interfaces for testing purposes. In this step, we define a `MockCustomerController` struct that embeds `mock.Mock`, which will allow us to set expectations and return values for each of the controller methods.

```
package router

import (
	"net/http"
	"github.com/stretchr/testify/mock"
)

type MockCustomerController struct {
	mock.Mock
}

func (cc *MockCustomerController) Add(w http.ResponseWriter, r *http.Request) {
	cc.Called(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"id": "cust105", "name": "arun", "email": "arun@gmail.com"}`))
}

func (cc *MockCustomerController) Get(w http.ResponseWriter, r *http.Request) {
	cc.Called(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write([]byte(`{"id": "cust105", "name": "arun", "email": "arun@gmail.com"}`))
}

func (cc *MockCustomerController) Update(w http.ResponseWriter, r *http.Request) {
	cc.Called(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write([]byte(`{"id": "cust105", "name": "arun", "email": "arun@gmail.com"}`))
}

func (cc *MockCustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	cc.Called(w, r)
	w.WriteHeader(http.StatusNoContent)
}

func (cc *MockCustomerController) GetAll(w http.ResponseWriter, r *http.Request) {
	cc.Called(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write([]byte(`[{"id": "cust105", "name": "arun", "email": "arun@gmail.com"}]`))
}

```

### Step 4: Setting Up Routes

Your `GetRouter` function defines the routes and ties them to controller methods. Here's an example:

```
package router

import (
	"net/http"
	"github.com/krishnakumarkp/customer-web/controller"
)

func GetRouter(controller controller.CustomerControllerInterface) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/customer", controller.Add)
	mux.HandleFunc("/customer/{id}", controller.Get)
	mux.HandleFunc("/customer", controller.GetAll)
	mux.HandleFunc("/customer/{id}", controller.Delete)
	mux.HandleFunc("/customer/{id}", controller.Update)
	return mux
}

```

his function defines routes for adding, getting, updating, and deleting customers.

### Step 5: Writing Tests with Mocked Controller

Next, let's write tests for the `GetRouter` function using the mocked controller. We will use `testify/mock` to mock the behavior of the controller methods and verify the routes are working as expected.

```
package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
				mockController.On("Add", mock.Anything, mock.Anything).Return()
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

```

### Explanation of the Test Code

1. **Test Structure**: Each test case defines the HTTP method, path, body, and expected response. The `mockSetup` function is used to specify the behavior of the mocked controller for that test case.
2. **Mock Expectations**: We use `mockController.On` to specify that when the controller methods are called with any arguments (`mock.Anything`), they should behave as defined in the mock methods.
3. **Request Execution**: The request is created using `httptest.NewRequest`, and the router is used to serve the request.
4. **Assertions**: We verify that the HTTP response status code and body match the expected values, and check that the mock expectations were met using `mockController.AssertExpectations`.