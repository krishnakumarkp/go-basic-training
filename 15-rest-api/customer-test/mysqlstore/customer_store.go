package mysqlstore

import (
	"database/sql"
	"errors"
	"go-training/customer-map/domain"
)

type CustomerStore struct {
	Store DataStore
}

func (cs CustomerStore) Create(customer domain.Customer) error {
	sqlStatement := "INSERT INTO customer(ID, Name, Email) VALUES (?,?,?)"
	insForm, err := cs.Store.Db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	_, err = insForm.Exec(customer.ID, customer.Name, customer.Email)
	return err
}

func (cs CustomerStore) GetById(id string) (domain.Customer, error) {
	var customer domain.Customer

	sqlStatement := "select id, name, email FROM customer where ID = ?"
	row := cs.Store.Db.QueryRow(sqlStatement, id)

	switch err := row.Scan(&customer.ID, &customer.Name, &customer.Email); err {
	case sql.ErrNoRows:
		err = errors.New("customer not found")
		return customer, err
	case nil:
		return customer, err
	default:
		panic(err)
	}
}
