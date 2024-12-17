package mysqlstore

import (
	"database/sql"
	"errors"

	"github.com/krishnakumarkp/customer-web/domain"
)

type CustomerStore struct {
	store DataStore
}

func NewCustomerStore(dbStore DataStore) CustomerStore {
	return CustomerStore{
		store: dbStore,
	}
}

func (cs CustomerStore) GetById(id string) (domain.Customer, error) {
	var customer domain.Customer

	sqlStatement := "select id, name, email FROM customer where ID = ?"
	row := cs.store.Db.QueryRow(sqlStatement, id)

	switch err := row.Scan(&customer.ID, &customer.Name, &customer.Email); err {
	case sql.ErrNoRows:
		err = errors.New("customer not found")
		return customer, err
	case nil:
		return customer, err
	default:
		return customer, err
	}
}
func (cs CustomerStore) GetAll() ([]domain.Customer, error) {
	var customers []domain.Customer
	var customer domain.Customer

	sqlStatement := `select id, name, email FROM customer ORDER BY id;`
	rows, err := cs.store.Db.Query(sqlStatement)
	if err != nil {
		return customers, errors.New("could not retrieve customers")
	}
	for rows.Next() {

		customer = domain.Customer{}
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (cs CustomerStore) Create(customer domain.Customer) error {
	sqlStatement := "INSERT INTO customer(ID, Name, Email) VALUES (?,?,?)"
	stmt, err := cs.store.Db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(customer.ID, customer.Name, customer.Email)
	return err
}

func (cs CustomerStore) Delete(id string) error {
	if id == "" {
		return errors.New("id is not passed")
	}
	sqlStatement := "DELETE FROM customer WHERE ID = ?"
	_, err := cs.store.Db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}

func (cs CustomerStore) Update(id string, customer domain.Customer) error {

	updateQuery := "UPDATE customer SET ID=?, Name=?, Email=? WHERE ID=?"
	_, err := cs.store.Db.Exec(updateQuery, customer.ID, customer.Name, customer.Email, id)

	if err != nil {
		return err
	}
	return nil
}
