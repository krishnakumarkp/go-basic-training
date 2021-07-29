package main

import (
	"dependencyinversion/customermanager/customer"
	"dependencyinversion/customermanager/mysqlstore"
	"fmt"
)

//Dependency Inversion Principle

// 1) High-level modules should not depend on low-level modules. Both should depend on abstractions.
// 2) Abstractions should not depend on details. Details should depend on abstractions.

type CustomerStore interface {
	Create(customer.Customer)
	Delete(string)
}

type CustomerController struct {
	Store CustomerStore
}

func (cc CustomerController) Add(c customer.Customer) {
	// cusomer controller is now dependent on interface CustomerStore in the same package.High-level modules is now not depend on low-level modules
	// module mysqlstore.Mysqlstore will have to impliment CustomerStore interface so it follows Dependency Inversion Principle

	// The design principle does not just change the direction of the dependency, as you might have expected when you read its
	// name for the first time. It splits the dependency between the high-level and low-level modules by introducing an
	// abstraction between them. So in the end, you get two dependencies:
	// 1) the high-level module depends on the abstraction, and
	// 2) the low-level depends on the same abstraction.

	// But now the interface is provided by the higher level component which the lover level componenet will have to impliment
	// so the dependency is now inverted.
	cc.Store.Create(c)
	fmt.Println("New Customer has been created")
}

func (cc CustomerController) Delete(id string) {
	cc.Store.Delete(id)
	fmt.Println("Customer has been deleted")
}

func main() {

	store := mysqlstore.Mysqlstore{}
	//here you can see we are passing MySql store to customer controller (dependency injection)
	controller := CustomerController{store}

	customer := customer.Customer{
		ID:    "cust105",
		Name:  "krishnakumar",
		Email: "krishna@test.com",
	}

	controller.Add(customer)
	controller.Delete("cust106")

}
