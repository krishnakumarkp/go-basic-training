package apputil

type Configuration struct {
	DBHost, DBPort, DBUser, DBPassword, Database string
}

var AppConfig Configuration

func init() {
	AppConfig = Configuration{}
}
