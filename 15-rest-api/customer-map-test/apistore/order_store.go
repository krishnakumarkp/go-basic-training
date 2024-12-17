package apistore

import (
	"bytes"
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

func (c OrderStore) CreateOrder(customerID string, order domain.Order) error {
	url := fmt.Sprintf("%s/customer/order/%s", c.BaseURL, customerID)
	payload, err := json.Marshal(order)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.New("failed to create order")
	}

	return nil
}

func (c OrderStore) DeleteOrdersByCustomerID(customerID string) error {
	url := fmt.Sprintf("%s/customer/order/%s", c.BaseURL, customerID)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return errors.New("failed to delete orders")
	}

	return nil
}
