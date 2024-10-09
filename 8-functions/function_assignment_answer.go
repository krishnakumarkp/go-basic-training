package main

import (
	"errors"
	"fmt"
)

//create a map names marks to store name and corresponding mark of students
var marks map[string]int

func init() {
	marks = make(map[string]int)
}

func add(k string, v int) error {
	if _, ok := marks[k]; ok {
		return errors.New(k + " is already present")
	}
	marks[k] = v
	return nil
}

func update(k string, v int) error {
	if _, ok := marks[k]; ok {
		marks[k] = v
		return nil
	}
	return errors.New(k + " is not found")
}

func remove(k string) error {
	if _, ok := marks[k]; ok {
		delete(marks, k)
		return nil
	}
	return errors.New(k + " is not found")
}

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
		fmt.Println("=========Vishal added=========")
	}

	if err := findAll(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("=========marks listed=========")
	}

	if err := update("Krishna", 85); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("=========krishna updated=========")
	}

	if err := findAll(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("=========marks listed=========")
	}

	if err := remove("Krishna"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("=========krishna deleted=========")
	}

	if err := findAll(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("=========marks listed=========")
	}
}
