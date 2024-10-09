package main

import "fmt"

//difference between Type Definition and Type alias

// Type Definitions
type myInt int
type myString string

// Type Alias Declarations
type myIntAlias = int
type myStringAlias = string

func main() {

	var name myString
	var anotherName string

	anotherName = "krishna"
	//name = anotherName
	name = myString(anotherName)
	fmt.Println(name)

	var location myStringAlias
	var anotherLocation string
	anotherLocation = "pune"

	location = anotherLocation
	fmt.Println(location)

}

// purpose of type aliases
// Refactoring
// The main purpose of introducing aliases was to make the refactoring of large codebases easier.
// Defining an alias from the old name to the new name allows the developers not to break
// any compatibility with existing clients.

//refer to alias-use folder

//package command
// Deprecated: Use github.com/docker/cli/cli/streams.In instead
//type InStream = streams.In

// Deprecated: Use github.com/docker/cli/cli/streams.Out instead
//type OutStream = streams.Out

//Readability
// The aliases can also be used to improve the readability of the code.
// Here is an example from the Go standard library and the disassembler package:
// type lookupFunc = func(addr uint64) (sym string, base uint64)
