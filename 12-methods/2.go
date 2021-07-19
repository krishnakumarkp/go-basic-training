package main

import (
	"fmt"
	"strings"
)

type name string

func (n name) Uppercase() string {
	return strings.ToUpper(string(n))
}

func main() {

	var myName name
	myName = "Krishnakumar"
	fmt.Println(myName.Uppercase())
}
