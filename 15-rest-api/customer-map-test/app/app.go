package app

import (
	"fmt"
	"go-training/customer-map/apistore"
	util "go-training/customer-map/apputil"
	"go-training/customer-map/controller"
	"go-training/customer-map/mysqlstore"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

func StartApp() {

	env := os.Getenv("APP_ENV")
	fmt.Println("env=", env)
	var configFile string
	var configPath string
	switch env {
	case "test":
		configFile = "test_app"
		configPath = "..config"
	default:
		configFile = "app"
		configPath = "config"
	}
	fmt.Println("config file =", configFile)
	viper.SetConfigName(configFile)
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	util.AppConfig.DBHost = viper.GetString("mysql.Host")
	util.AppConfig.DBPort = viper.GetString("mysql.Port")
	util.AppConfig.DBUser = viper.GetString("mysql.User")
	util.AppConfig.DBPassword = viper.GetString("mysql.Password")
	util.AppConfig.Database = viper.GetString("mysql.Database")
	util.AppConfig.OroderHost = viper.GetString("order.Host")

	config := mysqlstore.Config{
		Host:     util.AppConfig.DBHost,
		Port:     util.AppConfig.DBPort,
		User:     util.AppConfig.DBUser,
		Password: util.AppConfig.DBPassword,
		Database: util.AppConfig.Database,
	}

	dbStore, err := mysqlstore.NewDbStore(config)
	if err != nil {
		panic(err)
	}

	customerStore := mysqlstore.CustomerStore{Store: dbStore}
	apiStore := apistore.OrderStore{BaseURL: util.AppConfig.OroderHost}

	controller := controller.NewCustomerController(customerStore, apiStore)

	router := http.NewServeMux()

	//router.HandleFunc("GET /helloworld", controller.HelloHandler)

	router.HandleFunc("POST /customer", controller.Add)
	router.HandleFunc("GET /customer/{id}", controller.Get)
	router.HandleFunc("GET /customer/{id}/order", controller.GetOrders)
	http.ListenAndServe(":8081", router)

}

// func init() {
// 	viper.SetConfigName("app")
// 	viper.AddConfigPath("../config")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// 	util.AppConfig.DBHost = viper.GetString("mysql.Host")
// 	util.AppConfig.DBPort = viper.GetString("mysql.Port")
// 	util.AppConfig.DBUser = viper.GetString("mysql.User")
// 	util.AppConfig.DBPassword = viper.GetString("mysql.Password")
// 	util.AppConfig.Database = viper.GetString("mysql.Database")

// }
