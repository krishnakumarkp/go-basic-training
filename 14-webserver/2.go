// # Webserver

// The problem with ServeHTTP method of MyHttpHandler is that it responds to all the requests. But in a real world web app there will be differnet routes

// get http://localhost:8080/books

// get http://localhost:8080/books/213

// post http://localhost:8080/books

// to achive this we need a some thing called a multiplexer.

// tell about default multiplexer 8.go

// solution is ServeMux

// The ServeMux is a built-in struct type exported from the http package, that acts as HTTP request multiplexer.
// ServeMux implements the ServeHTTP method which will make it implement the Handler interface

package main

import (
	"fmt"
	"net/http"
)

type MyHelloWorldHandler struct{}

func (h MyHelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type MyHelloGolangHandler struct{}

func (h MyHelloGolangHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Golang!")
}

func main() {
	// create a new `ServeMux`
	mux := http.NewServeMux()

	handler1 := MyHelloWorldHandler{}

	handler2 := MyHelloGolangHandler{}

	// handle `/` route
	mux.Handle("/hello", handler1)
	mux.Handle("/hello/golang", handler2)

	// listen and serve using `ServeMux`
	http.ListenAndServe(":8080", mux)
}
