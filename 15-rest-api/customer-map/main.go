package main

import (
	"fmt"

	"github.com/kishnakumarkp/customermap/domain"
	"github.com/kishnakumarkp/customermap/mapstore"
)

type CustomerController struct {
	store domain.CustomerStore
}

func (cc CustomerController) Add(c domain.Customer) error {
	err := cc.store.Create(c)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("New Customer has been created")
	return err
}

func (cc CustomerController) Update(id string, c domain.Customer) error {
	err := cc.store.Update(id, c)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Customer has been updated")
	return err
}

func (cc CustomerController) Delete(id string) error {
	err := cc.store.Delete(id)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Customer has been deleted")
	return err
}

func (cc CustomerController) GetAll() error {
	customers, err := cc.store.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	for _, v := range customers {
		fmt.Println(v.ID, v.Name, v.Email)
	}
	return err
}

func main() {
	controller := CustomerController{
		store: mapstore.NewMapStore(),
	}

	customer := domain.Customer{
		ID:    "cust101",
		Name:  "krishnakumar",
		Email: "krishna@test.com",
	}

	controller.Add(customer)

	customer = domain.Customer{
		ID:    "cust102",
		Name:  "Devesh",
		Email: "Ddvesh@test.com",
	}

	controller.Add(customer)

	controller.GetAll()

	customer = domain.Customer{
		ID:    "cust101",
		Name:  "kk",
		Email: "kk@test.com",
	}

	controller.Update("cust101", customer)
	controller.Delete("cust101")
	controller.GetAll()

}
