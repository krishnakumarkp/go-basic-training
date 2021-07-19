package main

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func main() {

	// var h HandlerFunc

	// h = func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello Golang!")
	// }

	h := HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Golang!")
	})

	http.ListenAndServe(":80", h)

}
