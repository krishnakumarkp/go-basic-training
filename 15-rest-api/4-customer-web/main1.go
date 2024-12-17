package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	util "github.com/krishnakumarkp/customer-web/apputil"
	"github.com/krishnakumarkp/customer-web/controller"
	"github.com/krishnakumarkp/customer-web/mysqlstore"
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

	r := http.NewServeMux()

	r.HandleFunc("POST /customer", controller.Add)
	r.HandleFunc("GET /customer/{id}", controller.Get)
	r.HandleFunc("GET /customer", controller.GetAll)
	r.HandleFunc("DELETE /customer/{id}", controller.Delete)
	r.HandleFunc("PUT /customer/{id}", controller.Update)

	http.ListenAndServe(":8080", r)

}

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	util.AppConfig.DBHost = os.Getenv("MYSQL_HOST")
	util.AppConfig.DBPort = os.Getenv("MYSQL_PORT")
	util.AppConfig.DBUser = os.Getenv("MYSQL_USER")
	util.AppConfig.DBPassword = os.Getenv("MYSQL_PASSWORD")
	util.AppConfig.Database = os.Getenv("MYSQL_DATABASE")

}
