package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Creating an infoLog object for logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Register the helloWorldHandler function to handle requests to the root URL ("/").
	// Here, we use a closure to inject the infoLog dependency into the handler function.
	http.HandleFunc("/", getHelloWorldHandler(infoLog))

	// Start the web server on port 8080.
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// getHelloWorldHandler is a closure that returns a function to handle HTTP requests.
// It takes an infoLogger object as a parameter and returns a function that satisfies the http.HandlerFunc interface.
// This closure allows us to inject the infoLogger dependency into the handler function.
func getHelloWorldHandler(infoLogger *log.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Logging a message using the injected infoLogger.
		infoLogger.Printf("Requested hello world")

		// Writing "Hello, World!" to the response writer.
		fmt.Fprintf(w, "Hello, World!")
	}
}
