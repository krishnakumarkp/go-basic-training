package controller

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
	"go-training/customer/domain"
)

type CustomerController struct {
	Store domain.CustomerStore
}

func (cc CustomerController) Add(w http.ResponseWriter, r *http.Request) {
	var customer domain.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = cc.Store.Create(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jd, _ := json.Marshal(customer)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jd)
}

func (cc CustomerController) Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	customer, err := cc.Store.GetById(vars["id"])
	if err == nil {
		jd, _ := json.Marshal(customer)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		w.Write(jd)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (cc CustomerController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["id"]

	var customer domain.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err1 := cc.Store.GetById(customerId)
	if err1 != nil {
		http.Error(w, customerId, http.StatusNotFound)
		return
	}

	err = cc.Store.Update(customerId, customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jd, _ := json.Marshal(customer)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(jd)
}

//Delete removes the customer from store
func (cc CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := cc.Store.Delete(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

//GetAll gets all the cutomers
func (cc CustomerController) GetAll(w http.ResponseWriter, r *http.Request) {
	customers, err := cc.Store.GetAll()
	if err == nil {
		jd, _ := json.Marshal(customers)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		w.Write(jd)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
