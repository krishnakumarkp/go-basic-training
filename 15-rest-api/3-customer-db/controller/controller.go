package controller

import (
	"fmt"

	"github.com/krishnakumarkp/customer-db/domain"
)

type CustomerController struct {
	Store domain.CustomerStore
}

func (cc CustomerController) Add(c domain.Customer) {
	if err := cc.Store.Create(c); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("New Customer has been created")
}

func (cc CustomerController) Update(c domain.Customer) {
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
