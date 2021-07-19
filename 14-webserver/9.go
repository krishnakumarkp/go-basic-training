package main

import (
	"fmt"
	"log"
	"net/http"
)
//Middleware
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)
	logMux := logger(mux)
	http.ListenAndServe(":80", logMux)
}
