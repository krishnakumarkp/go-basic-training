package controller

import (
	"net/http"
)

// CustomerControllerInterface defines the methods for a customer controller.
type CustomerControllerInterface interface {
	Add(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}
