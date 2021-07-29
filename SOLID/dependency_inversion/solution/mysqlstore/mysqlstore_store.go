package mysqlstore

import (
	"dependencyinversion/customermanager/customer"
	"fmt"
)

type Mysqlstore struct {
}

func (ms Mysqlstore) Create(customer customer.Customer) {
	sqlStatement := "INSERT INTO customer(ID, Name, Email) VALUES (?,?,?)"
	fmt.Println(sqlStatement + " executed")
}

func (ms Mysqlstore) Delete(id string) {
	sqlStatement := "DELETE FROM customer WHERE ID = ?"
	fmt.Println(sqlStatement + " executed")
}
