package main

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

// MyHttpHandlerMux is a custom HTTP handler.
type MyHttpHandlerMux struct {
	routes map[string]HandlerFunc
}

// NewMyHttpHandlerMux creates a new instance of MyHttpHandlerMux.
func NewMyHttpHandlerMux() *MyHttpHandlerMux {
	return &MyHttpHandlerMux{
		routes: make(map[string]HandlerFunc),
	}
}

// ServeHTTP is the method required to implement the http.Handler interface.
func (h *MyHttpHandlerMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if there is a registered route for the request path
	if handlerFunc, ok := h.routes[r.URL.Path]; ok {
		// If yes, execute the associated handler function
		handlerFunc(w, r)
	} else {
		// If no, respond with a default message
		fmt.Fprintln(w, "Route not found")
	}
}

// HandleFunc registers a route with a specific handler function.
func (h *MyHttpHandlerMux) HandleFunc(route string, handler HandlerFunc) {
	h.routes[route] = handler
}

// Custom handler functions for specific routes
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, this is the custom hello route!")
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Goodbye, this is the custom goodbye route!")
}

func main() {
	// Create an instance of MyHttpHandlerMux
	myHandler := NewMyHttpHandlerMux()

	// Register custom handler functions for specific routes
	myHandler.HandleFunc("/hello", helloHandler)
	myHandler.HandleFunc("/goodbye", goodbyeHandler)

	// Start the server on port 8080
	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", myHandler)
}
