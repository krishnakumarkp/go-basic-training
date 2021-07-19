package main

import (
	"fmt"
)

//concreate type of interface variable
type Worker interface {
	Work()
}

type Person struct {
	name string
}

func (p Person) Work() {
	fmt.Println(p.name, " is working")
}

func describe(w Worker) {
	fmt.Printf("Concrete type %T value %v \n", w, w)
}

func main() {
	var w Worker
	w = Person{
		name: "Naveen",
	}
	//the concreate type of w is Person not worker
	describe(w)
	w.Work()
}
