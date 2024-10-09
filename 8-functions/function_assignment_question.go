package main

import (
	"errors"
	"fmt"
)

//Declare a map named marks to store name and corresponding mark of students
var marks map[string]int

// when creating applications in Go, if you need to be able to set up some form of state on the initial startup of your program.
// This could involve creating connections to databases, or loading in configuration from locally stored configuration files.
// Then you ise the init() functions. Init runs before main and is used only for initialisations.

func init() {
	//initialise marks
	marks = make(map[string]int)
}

//write function add to add name and mark to the marks map. return error name is already present if name exists in the map
func add(k string, v int) error {
	//your code goes here

	return errors.New(k + " is already present")

	return nil
}

//write function update which takes in name and mark and updates the marks map with new mark.
// returns error k + is not found if the student name is not present

//your code goes here

//write a function remove that takes in a name and removes it from maraks map. if name is not found
//return error key + "is not found"

//your code goes here

//write function findAll that prints all the entries in marks array
func findAll() error {

	if len(marks) == 0 {
		return errors.New("No values found")
	}
	for k, v := range marks {
		fmt.Printf("%s => %d \n", k, v)
	}
	return nil
}

func main() {
	if err := add("Krishna", 100); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("=========krishna added=========")
	}

	if err := add("Vishal", 90); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("=========vishal added=========")
	}
	//call findAll() and varify that krishna and Vishal are present in the map
	if err := findAll(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("=========marks listed=========")
	}

	//call function update to update krishna with new mark

	//call find all and verify that krishnas mark is updated
	if err := findAll(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("=========marks listed=========")
	}

	//call function remove to remove krishna from the marks map
	if err := remove("Krishna"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("=========krishna deleted=========")
	}
	//verify krishna is removed
	if err := findAll(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("=========marks listed=========")
	}

}
