package controller

import (
	"encoding/json"
	"net/http"

	"github.com/krishnakumarkp/customer-web/domain"
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
	id := r.PathValue("id")

	customer, err := cc.Store.GetById(id)
	if err != nil {
		if err.Error() == "customer not found" { // Assuming GetById returns this specific error message
			http.Error(w, "Customer not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	jd, _ := json.Marshal(customer)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Use 200 OK for a successful GET
	w.Write(jd)
}

func (cc CustomerController) Update(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("id")

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

// Delete removes the customer from store
func (cc CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	customerId := r.PathValue("id")

	// Check if the customer exists
	_, err := cc.Store.GetById(customerId)
	if err != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	// Proceed to delete the customer
	err = cc.Store.Delete(customerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// GetAll gets all the cutomers
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
