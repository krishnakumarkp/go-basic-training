package main

import (
	util "github.com/krishnakumarkp/customer-db/apputil"
	"github.com/krishnakumarkp/customer-db/controller"
	"github.com/krishnakumarkp/customer-db/domain"
	"github.com/krishnakumarkp/customer-db/mysqlstore"

	"github.com/spf13/viper"
)

func main() {

	config := mysqlstore.Config{
		Host:     util.AppConfig.DBHost,
		Port:     util.AppConfig.DBPort,
		User:     util.AppConfig.DBUser,
		Password: util.AppConfig.DBPassword,
		Database: util.AppConfig.Database,
	}
	// Creates a Mysql DB instance
	dbStore, err := mysqlstore.New(config)
	if err != nil {
		panic(err)
	}

	// Creates a customerstore which uses dbstore
	customerStore := mysqlstore.NewCustomerStore(dbStore)

	controller := controller.CustomerController{
		Store: customerStore,
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

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	util.AppConfig.DBHost = viper.GetString("mysql.Host")
	util.AppConfig.DBPort = viper.GetString("mysql.Port")
	util.AppConfig.DBUser = viper.GetString("mysql.User")
	util.AppConfig.DBPassword = viper.GetString("mysql.Password")
	util.AppConfig.Database = viper.GetString("mysql.Database")

}
