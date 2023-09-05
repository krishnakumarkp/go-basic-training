package mapstore

import (
	"errors"

	"github.com/krishnakumarkp/customer-map/domain"
)

type MapStore struct {
	store map[string]domain.Customer
}

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

func (m MapStore) Create(c domain.Customer) error {
	if _, ok := m.store[c.ID]; !ok {
		m.store[c.ID] = c
		return nil
	}
	return errors.New("Customer already present")
}

func (m MapStore) Update(id string, c domain.Customer) error {
	if _, ok := m.store[id]; ok {
		m.store[id] = c
		return nil
	}
	return errors.New("customer not present")
}

func (m MapStore) Delete(id string) error {
	if _, ok := m.store[id]; ok {
		delete(m.store, id)
		return nil
	}
	return errors.New("customer not present")
}

func (m MapStore) GetById(id string) (domain.Customer, error) {
	if c, ok := m.store[id]; ok {
		return c, nil
	}
	return domain.Customer{}, errors.New("customer not present")
}

func (m MapStore) GetAll() ([]domain.Customer, error) {
	var customers []domain.Customer
	err := errors.New("customers not present")

	for _, c := range m.store {
		customers = append(customers, c)
		err = nil
	}
	return customers, err

}
