package main

// The "net/http" library has methods to implement HTTP clients and servers
import (
	"fmt"
	"net/http"
)

// Go provides the net package https://golang.org/pkg/net/ that contains the utility packages to handle networking related stuff.
// This package contains the http package https://golang.org/pkg/net/http/ which is used to create HTTP servers and to make HTTP requests as a client

//The http package gives us some APIs to create a simple HTTP server.
//The ListenAndServe function

//func ListenAndServe(addr string, handler Handler) error

// This function starts an HTTP server and locks the process. It starts to listen to incoming HTTP requests and serve requests.
// Since this function locks the process, any code below this function call will not be executed.
// The process will be terminated if it returns an error or if the process is killed using a keyboard interrupt (using ctrl+c).

//Let’s understand the arguments of this function.

// The addr argument of type string is the address of the machine on which the server will be spawned.
// It is the combination of IP address and port which looks like "<ip-address>:<port>".
// We can use ":<port>" as well without an IP address in which case it will be reachable from all the addresses of the machine.

//The handle argument of the type Handler interface handles the incoming HTTP requests. The Handler interface looks like below.

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

//Let’s write the response handler of type Handler

type MyHttpHandler struct{}

func (h MyHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {

	handler := MyHttpHandler{}

	http.ListenAndServe(":8080", handler)
}
