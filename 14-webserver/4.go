package main

import (
	"fmt"
	"net/http"
)

// there is also another servemux http.DefaultServeMux which is a global ServeMux instance.
// if you  pass nil as the value of handler to the ListenAndServe function Go will internally use the http.DefaultServeMux

//We can add routes and handler functions to this instance by using the functions exposed by the http package.

//The http.HandleFunc function provided  does the exact same thing that the mux.HandleFunc does.
// It adds a handler function to the http.DefaultServeMux to handle HTTP requests of a specific route path.

func main() {
	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	// handle `/hello/golang` route to `http.DefaultServeMux`
	http.HandleFunc("/hello/golang", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Golang!")
	})

	// listen and serve using `http.DefaultServeMux`
	http.ListenAndServe(":8080", nil)

}
