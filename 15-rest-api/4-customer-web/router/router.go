package router

import (
	"net/http"

	"github.com/krishnakumarkp/customer-web/controller"
)

func GetRouter(controller controller.CustomerControllerInterface) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /customer", controller.Add)
	mux.HandleFunc("GET /customer/{id}", controller.Get)
	mux.HandleFunc("GET /customer", controller.GetAll)
	mux.HandleFunc("DELETE /customer/{id}", controller.Delete)
	mux.HandleFunc("PUT /customer/{id}", controller.Update)

	return mux
}
