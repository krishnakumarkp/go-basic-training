package apistore

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-training/customer-map/domain"
	"net/http"
)

type OrderStore struct {
	BaseURL string
}

func (c OrderStore) GetOrdersByCustomerID(customerID string) ([]domain.Order, error) {
	url := fmt.Sprintf("%s/customer/%s/order", c.BaseURL, customerID)
	fmt.Println(url, "called")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get orders")
	}

	var orders []domain.Order
	if err := json.NewDecoder(resp.Body).Decode(&orders); err != nil {
		return nil, err
	}

	return orders, nil
}
