package main

import "fmt"

//Anonymous struct
//An anonymous struct is a struct with no explicitly defined derived struct type.

func main() {
	krishna := struct {
		name            string
		age             int
		location, phone string
	}{
		name:     "Krishna",
		age:      45,
		location: "Pune",
		phone:    "8007344554",
	}

	fmt.Printf("Krishna = %+v\n", krishna)

}
