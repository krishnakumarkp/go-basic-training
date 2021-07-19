package main

import (
	"fmt"
	"log"
	"net/http"
)

//middle ware chaining
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func HelloGolangHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello Golang!")
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

func jsonHeaderSetter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusFound)
		h.ServeHTTP(w, r)
	})
}

func helloHeaderSetter(h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Hello", "World")
		h(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloHandler)
	mux.HandleFunc("/hello/golang", helloHeaderSetter(HelloGolangHandler))
	lh := logger(mux)
	hs := jsonHeaderSetter(lh)
	http.ListenAndServe(":80", hs)
}
