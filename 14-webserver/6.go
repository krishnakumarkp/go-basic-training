package main

import (
	"fmt"
	"net/http"
)

func main() {

	//same thing can be done using the build in  http.HandlerFunc

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Golang!")
	})

	http.ListenAndServe(":8080", h)

}

// HandleFunc is a convenient shortcut if you want to register that function on the DefaultServeMux
// HandlerFunc is needed, if you actually need an http.Handler, e.g. to pass it to some middleware.
