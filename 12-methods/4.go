package main

import (
	"fmt"
	"strings"
)

// Adding Methods to Function Types
// One of the interesting features of Go is its ability to take a function type
// and attach methods to it

type FullName func(string, string) string

func (f FullName) Uppercase(firstName string, lastName string) string {
	return strings.ToUpper(string(f(firstName, lastName)))
}

func main() {

	var myFullName FullName = func(fname string, lname string) string {
		return fname + " " + lname
	}

	fmt.Println(myFullName.Uppercase("krishna", "kumar"))

}
