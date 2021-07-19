package main

import "fmt"

//Pointer vs Value receiver

type Worker interface {
	Work()
}

type Person struct {
	name string
}

type Animal struct {
	name string
}

func (p Person) Work() {
	fmt.Println(p.name, " is working")
}

func (a *Animal) Work() {
	fmt.Println(a.name, " is working")
}

func main() {
	var w1 Worker

	w1 = Person{
		name: "Naveen",
	}
	w1.Work()

	var w2 Worker
	w2 = &Person{
		name: "Krishna",
	}
	w2.Work()

	var w3 Worker

	//The dynamic type of interface w3 is Animal and we can clearly see that
	// Animal does not implement the Work method but *Animal does.

	w3 = Animal{
		name: "horse",
	}

	w3.Work()

	var w4 Worker

	w4 = &Animal{
		name: "horse",
	}
	w4.Work()

}
