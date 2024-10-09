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

func HelloCybage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Cybage!")
}

func main() {
	// create a new `ServeMux`
	mux := http.NewServeMux()

	handler1 := MyHelloWorldHandler{}

	handler2 := MyHelloGolangHandler{}

	// handle `/` route
	mux.Handle("/hello", handler1)
	mux.Handle("/hello/golang", handler2)
	mux.HandleFunc("/hello/cybage", HelloCybage)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is my webserver root!")
	})

	// listen and serve using `ServeMux`
	http.ListenAndServe(":8080", mux)
}
