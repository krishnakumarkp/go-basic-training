package domain

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomerOrder struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Orders []Order `json:"orders"`
}

type CustomerStore interface {
	Create(Customer) error
	// Update(string, Customer) error
	// Delete(string) error
	GetById(string) (Customer, error)
	// GetAll() ([]Customer, error)
}

type Order struct {
	BookID   string `json:"book_id"`
	Quantity int    `json:"quantity"`
}

// OrderStore represents an interface for interacting with the order  API
type OrderStore interface {
	GetOrdersByCustomerID(customerID string) ([]Order, error)
	//CreateOrder(customerID string, order Order) error
	//DeleteOrdersByCustomerID(customerID string) error
}
