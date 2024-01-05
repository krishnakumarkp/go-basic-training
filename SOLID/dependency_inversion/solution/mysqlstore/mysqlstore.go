package mysqlstore

import (
	"dependencyinversion/customermanager/customer"
	"fmt"
)

type Mysqlstore struct {
}

func (ms Mysqlstore) Create(customer customer.Customer) {
	sqlStatement := "INSERT INTO customer(ID, Name, Email) VALUES (%s,%s,%s)"
	fmt.Printf(sqlStatement+" executed \n", customer.ID, customer.Name, customer.Email)
}

func (ms Mysqlstore) Delete(id string) {
	sqlStatement := "DELETE FROM customer WHERE ID = %s"
	fmt.Printf(sqlStatement+" executed\n", id)
}
