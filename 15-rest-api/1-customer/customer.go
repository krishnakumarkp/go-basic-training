package main

import (
	"errors"
	"fmt"
)

type Customer struct {
	ID, Name, Email string
}
type CustomerStore interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetById(string) (Customer, error)
	GetAll() ([]Customer, error)
}

type MapStore struct {
	store map[string]Customer
}

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]Customer)}
}

func (m MapStore) Create(c Customer) error {
	if _, ok := m.store[c.ID]; !ok {
		m.store[c.ID] = c
		return nil
	}
	return errors.New("Customer already present")
}

func (m MapStore) Update(id string, c Customer) error {
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

func (m MapStore) GetById(id string) (Customer, error) {
	if c, ok := m.store[id]; ok {
		return c, nil
	}
	return Customer{}, errors.New("customer not present")
}

func (m MapStore) GetAll() ([]Customer, error) {
	var customers []Customer
	err := errors.New("customers not present")

	for _, c := range m.store {
		customers = append(customers, c)
		err = nil
	}
	return customers, err

}

type CustomerController struct {
	Store CustomerStore
}

func (cc CustomerController) Add(c Customer) {
	if err := cc.Store.Create(c); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("New Customer has been created")
}

func (cc CustomerController) Update(c Customer) {
	if err := cc.Store.Update(c.ID, c); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Customer has been updated")
}

func (cc CustomerController) Delete(id string) {
	if err := cc.Store.Delete(id); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Customer has been Deleted")
}

func (cc CustomerController) Get(id string) {
	c, err := cc.Store.GetById(id)
	if err == nil {
		fmt.Println(c)
		return
	}
	fmt.Println("Error: ", err)
}

func (cc CustomerController) GetAll() {
	customers, err := cc.Store.GetAll()
	if err == nil {
		fmt.Println(customers)
	} else {
		fmt.Println("Error: ", err)
	}
}

func main() {
	controller := CustomerController{
		Store: NewMapStore(),
	}

	customer1 := Customer{
		ID:    "cust101",
		Name:  "Krishnakumar",
		Email: "krishnakumarkp@gmail.com",
	}

	customer2 := Customer{
		ID:    "cust101",
		Name:  "Krishnakumar",
		Email: "krishnakumarkp@gmail.com",
	}

	customer3 := Customer{
		ID:    "cust101",
		Name:  "krishna",
		Email: "krishna@gmail.com",
	}

	customer4 := Customer{
		ID:    "cust102",
		Name:  "parvathy",
		Email: "parvathy@gmail.com",
	}

	controller.Add(customer1)
	controller.Add(customer2)
	controller.Add(customer4)
	controller.Update(customer3)
	controller.Get("cust101")

	controller.GetAll()

	controller.Delete("cust101")
	controller.Delete("cust101")
	controller.Get("cust101")

	controller.GetAll()
}
