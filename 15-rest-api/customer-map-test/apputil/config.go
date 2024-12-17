package apputil

type Configuration struct {
	DBHost, DBPort, DBUser, DBPassword, Database, OroderHost string
}

var AppConfig Configuration

func init() {
	AppConfig = Configuration{}
}
