

### Mocking Interfaces with `golang/mock`

This tutorial will guide you through the process of using the `golang/mock` package to generate mocks for interfaces, which is useful for unit testing your Go applications.

#### Step 1: Install Dependencies

First, you need to install `golang/mock`. You can use the following command to install the package:

```
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen@latest

export PATH=$PATH:$(go env GOPATH)/bin
```

#### Step 2: Define the Interface to Mock

In our case, you have a `CustomerStore` interface in the `domain` package that you want to mock for testing.

#### Step 3: Generate Mocks with `mockgen`

To generate mocks for the `CustomerStore` interface, you need to run the `mockgen` tool.

```
mockgen -source=domain/domain.go -destination=mocks/mock_customer_store.go -package=mocks CustomerStore
```

Explanation of the command:

- `-source`: Path to the file that contains the interface you want to mock.
- `-destination`: Path to the file where the generated mock code will be written.
- `-package`: The name of the package where the mock will reside (in this case, `mocks`).

After running this command, you should have a new file called `mock_customer_store.go` in the `mocks` directory.

#### Step 4: Use the Mock in Your Tests



Now that we have the mock implementation, you can use it in your tests. Here’s an example of how to use the generated mock in a test for the `Add` method in the `CustomerController`.

```
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

	// Test case 3: Store Create Error
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

```

In this test:

1. We create a new `gomock.Controller` which will manage the lifecycle of mock objects.
2. We use `mocks.NewMockCustomerStore(ctrl)` to create a mock of the `CustomerStore` interface.
3. We define expectations on the mock using `mockStore.EXPECT()`. This specifies what method we expect to be called and what it should return.
4. We test different cases such as valid data, invalid JSON, and error handling.