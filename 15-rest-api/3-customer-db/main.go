package main

import (
	"fmt"

	util "go-training/customer/apputil"
	"go-training/customer/domain"
	"go-training/customer/mysqlstore"

	"github.com/spf13/viper"
)

type CustomerController struct {
	store domain.CustomerStore
}

func (cc CustomerController) Add(c domain.Customer) error {
	err := cc.store.Create(c)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("New Customer has been created")
	return err
}

func (cc CustomerController) Update(id string, c domain.Customer) error {
	err := cc.store.Update(id, c)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Customer has been updated")
	return err
}

func (cc CustomerController) Delete(id string) error {
	err := cc.store.Delete(id)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Customer has been deleted")
	return err
}

func (cc CustomerController) GetAll() error {
	customers, err := cc.store.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Get all")
	for _, v := range customers {
		fmt.Println(v.ID, v.Name, v.Email)
	}
	return err
}

func (cc CustomerController) Get(id string) error {

	customer, err := cc.store.GetById(id)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	} else {
		fmt.Println("Get customer ", id)
		fmt.Println(customer.ID, customer.Name, customer.Email)
	}
	return err
}

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

	controller := CustomerController{
		store: customerStore,
	}

	customer := domain.Customer{
		ID:    "cust105",
		Name:  "krishnakumar",
		Email: "krishna@test.com",
	}

	controller.Add(customer)

	controller.Get("cust105")

	customer = domain.Customer{
		ID:    "cust106",
		Name:  "Devesh",
		Email: "Ddvesh@test.com",
	}

	controller.Add(customer)

	controller.GetAll()

	customer = domain.Customer{
		ID:    "cust106",
		Name:  "kk",
		Email: "kk@test.com",
	}

	controller.Update("cust106", customer)
	controller.Delete("cust106")
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
