package app

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/krishnakumarkp/customer-web/apputil"
	"github.com/krishnakumarkp/customer-web/controller"
	"github.com/krishnakumarkp/customer-web/mysqlstore"
	"github.com/krishnakumarkp/customer-web/router"
)

func StartApp() {

	env := os.Getenv("APP_ENV")
	envfile := ".env"
	if env == "test" {
		envfile = "../../.env.test" // Default to development
	}

	if err := godotenv.Load(envfile); err != nil {
		panic(err)
	}
	apputil.AppConfig.DBHost = os.Getenv("MYSQL_HOST")
	apputil.AppConfig.DBPort = os.Getenv("MYSQL_PORT")
	apputil.AppConfig.DBUser = os.Getenv("MYSQL_USER")
	apputil.AppConfig.DBPassword = os.Getenv("MYSQL_PASSWORD")
	apputil.AppConfig.Database = os.Getenv("MYSQL_DATABASE")

	config := mysqlstore.Config{
		Host:     apputil.AppConfig.DBHost,
		Port:     apputil.AppConfig.DBPort,
		User:     apputil.AppConfig.DBUser,
		Password: apputil.AppConfig.DBPassword,
		Database: apputil.AppConfig.Database,
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

	r := router.GetRouter(controller)

	http.ListenAndServe(":8080", r)

}
