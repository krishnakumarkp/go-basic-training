package main

import (
	"singleresponsibility/library/book"
	"singleresponsibility/library/persistence"
)

//Single Responsibility Principle
//A class should have only one reason to change.

func main() {

	book := book.Book{}
	book.Title = "The Go Programming Language"
	book.Content = "The Go Programming Language is the authoritative resource for any programmer who wants to learn Go"

	persistence := persistence.Persistence{}
	persistence.Save(book)
	//Could I “Save” the book in a file, or in a DataBase?
	// Yes i can do it without changing the book struct, I need to change only persistence struct

	// The Single Responsibility Principle encourages you to structure the functions,
	// types, and methods into packages that exhibit natural cohesion; the types
	// belong together, the functions serve a single purpose.
}
