package main

import (
	"net/http"

	util "github.com/krishnakumarkp/customer-web/apputil"
	"github.com/krishnakumarkp/customer-web/controller"
	"github.com/krishnakumarkp/customer-web/mysqlstore"

	"github.com/gorilla/mux"
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

	router := mux.NewRouter()

	router.HandleFunc("/customer", controller.Add).Methods("POST")
	router.HandleFunc("/customer/{id}", controller.Get).Methods("GET")
	router.HandleFunc("/customer", controller.GetAll).Methods("GET")
	router.HandleFunc("/customer/{id}", controller.Delete).Methods("DELETE")
	router.HandleFunc("/customer/{id}", controller.Update).Methods("PUT")

	http.ListenAndServe(":8080", router)

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
