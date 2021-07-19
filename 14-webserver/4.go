package main

import (
	"fmt"
	"net/http"
)

type MyHttpHandler struct{}

func (h MyHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// The ServeMux is a built-in struct type exported from the http package,
// that acts as HTTP request multiplexer.
// ServeMux implements the ServeHTTP method which will make it implement
// the Handler interfac

func main() {
	// create a new `ServeMux`
	mux := http.NewServeMux()

	handler := MyHttpHandler{}

	// handle `/` route
	mux.Handle("/", handler)

	// handle `/hello/golang` route
	mux.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	// listen and serve using `ServeMux`
	http.ListenAndServe(":8080", mux)
}
