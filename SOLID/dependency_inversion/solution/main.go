package main

import (
	"dependencyinversion/customermanager/customer"
	"dependencyinversion/customermanager/mysqlstore"
	"fmt"
)

//Dependency Inversion Principle

// 1) High-level modules should not depend on low-level modules. Both should depend on abstractions.
// 2) Abstractions should not depend on details. Details should depend on abstractions.

type CustomerController struct {
	Store mysqlstore.Mysqlstore
}

func (cc CustomerController) Add(c customer.Customer) {
	// cusomer controller is dependent on Mysqlstore.High-level modules depend on low-level modules which is
	//violation of Dependency Inversion Principle
	cc.Store.Create(c)
	fmt.Println("New Customer has been created")
}

func (cc CustomerController) Delete(id string) {
	cc.Store.Delete(id)
	fmt.Println("Customer has been deleted")
}

func main() {

	store := mysqlstore.Mysqlstore{}
	//here you can see we are passing Mysqlstore to customer controller (dependency injection)
	controller := CustomerController{store}

	customer := customer.Customer{
		ID:    "cust105",
		Name:  "krishnakumar",
		Email: "krishna@test.com",
	}

	controller.Add(customer)
	controller.Delete("cust106")

}
