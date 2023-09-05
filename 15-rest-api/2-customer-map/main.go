package main

import (
	"github.com/krishnakumarkp/customer-map/controller"
	"github.com/krishnakumarkp/customer-map/domain"
	"github.com/krishnakumarkp/customer-map/mapstore"
)

func main() {
	controller := controller.CustomerController{
		Store: mapstore.NewMapStore(),
	}

	customer1 := domain.Customer{
		ID:    "cust101",
		Name:  "Krishnakumar",
		Email: "krishnakumarkp@gmail.com",
	}

	customer2 := domain.Customer{
		ID:    "cust101",
		Name:  "Krishnakumar",
		Email: "krishnakumarkp@gmail.com",
	}

	customer3 := domain.Customer{
		ID:    "cust101",
		Name:  "krishna",
		Email: "krishna@gmail.com",
	}

	customer4 := domain.Customer{
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
