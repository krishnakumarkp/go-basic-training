package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// handle `/` route
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	// // handle `/hello/golang` route
	// router.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
	// 	name := req.URL.Query().Get("name")
	// 	fmt.Fprint(res, "Hello "+name)
	// })

	router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	}).Methods("GET")

	// listen and serve using `GorillaMux`
	http.ListenAndServe(":8080", router)

	//http://localhost:8080/books/golang/page/123

}
