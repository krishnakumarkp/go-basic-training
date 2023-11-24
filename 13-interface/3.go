package main

import "fmt"

type Worker interface {
	Work()
}

type Person struct {
	name string
}

func (p Person) Work() {
	fmt.Println(p.name, " is working")
}

func main() {
	//A variable of an interface type
	var w Worker

	//zero value and type of the interface is nil
	if w == nil {
		fmt.Printf("w is nil and has type %T value %v\n", w, w)
	}

	w = Person{
		name: "Naveen",
	}

	//static type of w is worker but
	//the concreate(dynamic) type of w is Person not worker

	fmt.Printf("Concrete type %T value %v \n", w, w)

	w.Work()

}
