package main

import "singleresponsibility/library/book"

//Single Responsibility Principle
//A class should have only one reason to change.

func main() {

	book := book.Book{}
	book.Title = "The Go Programming Language"
	book.Content = "The Go Programming Language is the authoritative resource for any programmer who wants to learn Go"
	book.Save()
	//Could I “Save” the book in a file, or in a DataBase?
	// Yes but the struct book should be change for no reason.
}
