package persistence

import (
	"fmt"
	"singleresponsibility/library/book"
)

type Persistence struct {
}

func (p Persistence) Save(b book.Book) {
	fmt.Println("Saving in the Terminal .... ")
	fmt.Println(b.GetTitle())
	fmt.Println(b.GetContent())
}
