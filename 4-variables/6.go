package main

import "fmt"

//Global Variable(Declared outside a block or a function)
var company string = "cybage"

func main() {
	//Local Variable (Declared inside a block or a function)
	var name string = "krishna"
	fmt.Println(name)
	if company == "cybage" {
		var name = company
		fmt.Println(name)
	}
	fmt.Println(name)
	PrintCompany()
}
func PrintCompany() {
	//local variable with the same name as that of grloval variable
	//variable shadowing
	var company = "google"
	fmt.Println(company)
}
