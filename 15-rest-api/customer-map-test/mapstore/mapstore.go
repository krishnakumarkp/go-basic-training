package mapstore

import (
	"errors"
	"go-training/customer-map/domain"
)

type MapStore struct {
	Store map[string]domain.Customer
}

func NewMapStore() *MapStore {
	return &MapStore{Store: make(map[string]domain.Customer)}
}

func (m MapStore) Create(c domain.Customer) error {
	if _, ok := m.Store[c.ID]; !ok {
		m.Store[c.ID] = c
		return nil
	}
	return errors.New("Customer already present")
}

func (m MapStore) GetById(id string) (domain.Customer, error) {
	if c, ok := m.Store[id]; ok {
		return c, nil
	}
	return domain.Customer{}, errors.New("customer not present")
}
