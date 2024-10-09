package controller

import (
	"encoding/json"
	"go-training/customer-map/domain"
	"net/http"
)

type CustomerController struct {
	Store    domain.CustomerStore
	ApiStore domain.OrderStore
}

func NewCustomerController(store domain.CustomerStore, apistore domain.OrderStore) CustomerController {
	return CustomerController{
		Store:    store,
		ApiStore: apistore,
	}
}

func (cc CustomerController) Get(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	customer, err := cc.Store.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		jd, err := json.Marshal(customer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jd)
	}
}

func (cc CustomerController) Add(w http.ResponseWriter, r *http.Request) {
	var customer domain.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := cc.Store.Create(customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jd, err := json.Marshal(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jd)
}

func (cc CustomerController) GetOrders(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	customer, err := cc.Store.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	orders, err := cc.ApiStore.GetOrdersByCustomerID(customer.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var customerOrdrs = domain.CustomerOrder{
		ID:     customer.ID,
		Name:   customer.Name,
		Email:  customer.Email,
		Orders: orders,
	}
	w.Header().Set("Content-Type", "application/json")
	jd, err := json.Marshal(customerOrdrs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jd)
}
