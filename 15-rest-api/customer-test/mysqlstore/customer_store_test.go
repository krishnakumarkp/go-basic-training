package mysqlstore

import (
	"errors"
	"os"
	"testing"

	util "go-training/customer-map/apputil"
	"go-training/customer-map/domain"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var dbStore DataStore

func TestMain(m *testing.M) {

	viper.SetConfigName("app_test")
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

	exitVal := m.Run()

	//after test

	os.Exit(exitVal)

}
func TestMySqlstore_Create(t *testing.T) {

	customerStore := CustomerStore{Store: dbStore}

	customer := domain.Customer{
		ID:    "cust101",
		Name:  "Krishnakumar",
		Email: "krishnakumarkp@gmail.com",
	}

	t.Run("Create a new customer", func(t *testing.T) {
		err := customerStore.Create(customer)
		assert.NoError(t, err, "Expected no error")
	})

	t.Run("Create a duplicate customer", func(t *testing.T) {
		// Try to create a duplicate customer
		err := customerStore.Create(customer)
		expectedErr := errors.New("Error 1062 (23000): Duplicate entry 'cust101' for key 'customer.ID'")
		assert.EqualError(t, err, expectedErr.Error(), "Expected error Error 1062 (23000): Duplicate entry 'cust101' for key 'customer.ID'")

	})

}

func truncateTable(tableName string, dbStore DataStore) error {
	// Construct the SQL statement to truncate the table
	truncateQuery := "TRUNCATE TABLE " + tableName

	// Execute the SQL statement
	_, err := dbStore.Db.Exec(truncateQuery)
	return err
}
