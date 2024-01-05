package mongostore

import (
	"dependencyinversion/customermanager/customer"
	"fmt"
)

type Mongostore struct {
}

func (ms Mongostore) Create(customer customer.Customer) {
	documentQuery := `db.customer.insertOne({
		ID: %s,
		Name: %s,
		Email: %s
	  });`

	fmt.Printf(documentQuery+" executed\n", customer.ID, customer.Name, customer.Email)
}

func (ms Mongostore) Delete(id string) {
	documentQuery := "db.customer.deleteOne({ ID: " + id + " });"
	fmt.Println(documentQuery + " executed")
}
