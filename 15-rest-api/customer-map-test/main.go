package main

import "go-training/customer-map/app"

func main() {

	app.StartApp()

}

// func main() {

// 	config := mysqlstore.Config{
// 		Host:     util.AppConfig.DBHost,
// 		Port:     util.AppConfig.DBPort,
// 		User:     util.AppConfig.DBUser,
// 		Password: util.AppConfig.DBPassword,
// 		Database: util.AppConfig.Database,
// 	}

// 	dbStore, err := mysqlstore.NewDbStore(config)
// 	if err != nil {
// 		panic(err)
// 	}

// 	customerStore := mysqlstore.CustomerStore{Store: dbStore}

// 	controller := controller.NewCustomerController(customerStore)

// 	router := http.NewServeMux()

// 	//router.HandleFunc("GET /helloworld", controller.HelloHandler)

// 	router.HandleFunc("POST /customer", controller.Add)
// 	router.HandleFunc("GET /customer/{id}", controller.Get)
// 	http.ListenAndServe(":8081", router)

// }

// func init() {
// 	viper.SetConfigName("app")
// 	viper.AddConfigPath("config")
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
