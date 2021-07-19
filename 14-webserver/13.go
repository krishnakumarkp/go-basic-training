package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Customer struct to hold all cutomer details
type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Add(w http.ResponseWriter, r *http.Request) {
	var customer Customer

	err := json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jd, _ := json.Marshal(customer)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jd)
}

func Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	customer := Customer{
		ID:    vars["id"],
		Name:  "krishna",
		Email: "krishna@cybage.com",
	}
	jd, _ := json.Marshal(customer)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(jd)
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/customer", Add).Methods("POST")
	router.HandleFunc("/customer/{id}", Get).Methods("GET")
	//router.Use(logger)
	lrouter := logger(router)
	http.ListenAndServe(":8080", lrouter)
}
