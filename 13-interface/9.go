package main

import (
	"fmt"
)

// Animal is an interface with a method Speak.
type Animal interface {
	Speak() string
}

// Dog is a type that implements the Animal interface.
type Dog struct{}

// Speak is a method of Dog that returns a string.
func (d Dog) Speak() string {
	return "Woof!"
}

// Cat is a type that implements the Animal interface.
type Cat struct{}

// Speak is a method of Cat that returns a string.
func (c Cat) Speak() string {
	return "Meow!"
}

// PrintAnimalSound is a function that takes an animals and prints its sounds using type assertion.
func PrintAnimalSound(animal Animal) {

	// 1. for some reason we need to get concrete type of dog in side the function
	// The type assertion x.(T) asserts that the concrete value stored in x is of type T,
	// and that x is not nil.

	// If T is not an interface, it asserts that the dynamic type of x is identical to T.
	// If T is an interface, it asserts that the dynamic type of x implements T.
	dog := animal.(Dog) //this will panic if animal is not dog
	fmt.Println("This is a dog:", dog.Speak())

	//2. this will return false and will not panic
	if dog, ok := animal.(Dog); ok {
		fmt.Println("This is a dog:", dog.Speak())
	}

	//3. we can make it better by if else

	if dog, ok := animal.(Dog); ok {
		fmt.Println("This is a dog:", dog.Speak())
	} else if cat, ok := animal.(Cat); ok {
		fmt.Println("This is a cat:", cat.Speak())
	} else {
		fmt.Println("Unknown animal type!")
	}

	//4. it can be still made better by using switch case

	// We can find out the underlying dynamic value of an interface using the syntax i.(Type)
	// where i is a variable of type interface and Type is a type that implements the interface.

	// Type switch
	switch v := animal.(type) { //this syntax will only work in switch statement.
	case Dog:
		fmt.Println("This is a dog:", v.Speak())
	case Cat:
		fmt.Println("This is a cat:", v.Speak())
	default:
		fmt.Println("Unknown animal type!")
	}

}

func main() {
	// Creating  Dog and Cat instances.

	dog := Dog{}
	cat := Cat{}

	// Calling the function to print animal sounds.
	PrintAnimalSound(dog)
	PrintAnimalSound(cat)
}
