package main

import "fmt"

func main() {
	fmt.Println(getName())
}

//If the deferred function is a function literal and the surrounding
//function has named result parameters that are in scope within the literal,
//the deferred function may access and modify the result parameters before they
//are returned

func getName() (n string) {
	defer func() {
		n = "golang"
	}()
	n = "python"
	return
}
