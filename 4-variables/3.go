package main

//Different ways of Declaring and initializing  Variablses
func main() {
	var name string
	name = krishna

	var name string = "krishna"
	var name, anotherName string = "krishna", "anothername"

	//type inference
	var name = "krishna"

	//short variable declaration
	name := "krishna"

	//var grouped
	var (
		name     string = "krishna"
		age      int    = 32
		location string = "pune"
	)

}
