package mysqlstore

import (
	"errors"
	"go-training/customer-map/domain"
	"os"
	"testing"

	util "go-training/customer-map/apputil"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var dbStore DataStore

func TestMain(m *testing.M) {
	viper.SetConfigName("test_app")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	util.AppConfig.DBHost = viper.GetString("mysql.Host")
	util.AppConfig.DBPort = viper.GetString("mysql.Port")
	util.AppConfig.DBUser = viper.GetString("mysql.User")
	util.AppConfig.DBPassword = viper.GetString("mysql.Password")
	util.AppConfig.Database = viper.GetString("mysql.Database")

	config := Config{
		Host:     util.AppConfig.DBHost,
		Port:     util.AppConfig.DBPort,
		User:     util.AppConfig.DBUser,
		Password: util.AppConfig.DBPassword,
		Database: util.AppConfig.Database,
	}

	// Creates a Mysql DB instance
	dbStore, err = NewDbStore(config)

	if err != nil {
		panic(err)
	}

	defer dbStore.Db.Close()

	// Truncate the table before running tests
	tableName := "customer"
	if err = truncateTable(tableName, dbStore); err != nil {
		panic(err)
	}

	// Run the tests
	exitVal := m.Run()

	os.Exit(exitVal)

}

func TestMapStore_Create(t *testing.T) {

	customerStore := CustomerStore{Store: dbStore}

	customer1 := domain.Customer{
		ID:    "cust101",
		Name:  "Krishnakumar",
		Email: "krishnakumarkp@gmail.com",
	}

	t.Run("Create a new customer", func(t *testing.T) {
		err := customerStore.Create(customer1)
		assert.NoError(t, err, "Expected no error")
	})

	t.Run("Create a duplicate customer", func(t *testing.T) {
		// Try to create a duplicate customer
		err := customerStore.Create(customer1)
		expectedErr := errors.New("Error 1062 (23000): Duplicate entry 'cust101' for key 'customer.ID'")
		assert.EqualError(t, err, expectedErr.Error(), "Expected error Error 1062 (23000): Duplicate entry 'cust101' for key 'customer.ID'")

	})

}

func TestMapStore_GetById(t *testing.T) {

	customerStore := CustomerStore{Store: dbStore}

	// Create a customer
	customerStore.Create(domain.Customer{ID: "1", Name: "John", Email: "john@test.com"})

	// Get the customer by ID
	c, err := customerStore.GetById("1")
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, "1", c.ID, "Expected customer ID to be '1'")
	assert.Equal(t, "John", c.Name, "Expected customer name to be 'John'")

	// Try to get a non-existent customer
	_, err = customerStore.GetById("2")
	expectedErr := errors.New("customer not found")
	assert.EqualError(t, err, expectedErr.Error(), "Expected error 'customer not found'")
}

func truncateTable(tableName string, dbStore DataStore) error {
	// Construct the SQL statement to truncate the table
	truncateQuery := "TRUNCATE TABLE " + tableName

	// Execute the SQL statement
	_, err := dbStore.Db.Exec(truncateQuery)
	return err
}
