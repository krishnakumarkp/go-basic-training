package main

import (
	"net/http"
)

func main() {

	// 1) create a custom struct type MyHandler and impliment http.Handler interface and pass in variable handler of type MyHandler.
	// to ListenAndServe methd.
	// handler := MyHandler{}
	//this handler should show messge "Hello Cybage!" for the request http://localhost:8080/

	// 2) create function helloGoogle(res http.ResponseWriter, req *http.Request)  and convert it to type type http.HandlerFunc and pass that as handler
	// this handler should show messge "Hello Google from function !" for the request http://localhost:8080/

	// 3) create handler variable of any name of type *http.ServeMux  using the function http.NewServeMux() and attach the MyHandler you created in step 1
	// to servemux for the route "/hello/cybage" also attach function helloGoogle to servemux for the route "/hello/google" and pass the servemux to listen and serve.
	// after this step  http://localhost:8080/hello/cybage will show "Hello Cybage!" and  http://localhost:8080/hello/google will show  "Hello Google from function !"

	http.ListenAndServe(":8080", handler)

}
