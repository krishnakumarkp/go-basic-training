package main

//https://gocodecloud.com/blog/2016/11/15/simple-golang-http-request-context-example/

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		username := ctx.Value("Username")

		fmt.Fprintf(w, "hello %s \n", username)
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)
		ctx := context.WithValue(r.Context(), "Username", "krishnakumar")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", hello)
	contextmux := AddContext(mux)
	http.ListenAndServe(":8080", contextmux)
}
