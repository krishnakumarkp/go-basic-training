package app

import (
	"go-training/customer-map/mysqlstore"
	"go-training/customer-map/router"
	"net/http"
	"os"

	util "go-training/customer-map/apputil"

	"github.com/spf13/viper"
)

func StartApp() {

	env := os.Getenv("APP_ENV")

	var configFile string
	var configPath string

	switch env {
	case "test":
		configFile = "app_test"
		configPath = "../config"
	default:
		configFile = "app"
		configPath = "config"
	}

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
	router := router.GetRouter(customerStore)

	http.ListenAndServe(":8081", router)

}
