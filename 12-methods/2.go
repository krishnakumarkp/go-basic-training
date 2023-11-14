package main

import (
	"fmt"
	"strings"
)

type name string

// methods can be created only for defined types
// cannot define methods for pointer types type name *string

// method
func (n name) Uppercase() string {
	return strings.ToUpper(string(n))
}

// function
func Uppercase(n name) string {
	return strings.ToUpper(string(n))
}

type newNameType name

func main() {

	var myName name

	myName = "Krishnakumar"

	// function call
	fmt.Println(Uppercase(myName))

	// method call
	fmt.Println(myName.Uppercase())

	// method can be called on value of the type or direclty
	// on the type and pass on the value as first parameter

	//called on value of the type
	fmt.Println(myName.Uppercase())

	//called direclty on the type and pass on the value
	fmt.Println(name.Uppercase(myName))

	//you can see the function syntax by printing the type

	//by type
	fmt.Printf("%T\n", name.Uppercase)

	//value
	fmt.Printf("%T\n", myName.Uppercase)

	//even if the method receiver is a pointer type the value need not be
	//a pointer go will manage it internally

	// func (n *name) Uppercase() string {
	// 	return strings.ToUpper(string(*n))
	// }

	// var myName name
	// myName = "Krishnakumar"

	// var myNewName *name

	// myNewName = &myName

	// fmt.Println(myNewName.Uppercase())

	// myOtherName := new(name)
	// *myOtherName = "rahul"

	// fmt.Println((*myOtherName).Uppercase())

	//type created from another type will not obtain the method defined on source type

	var newName newNameType

	newName = "sheetal"

	fmt.Println(newName.Uppercase())

}
