package main

// Type assertions

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

// PrintAnimalSound is a function that takes an animal
// inside this function we dont knwo if the animal is a dog or a cat
func PrintAnimalSound(animal Animal) {

	fmt.Println(animal.Speak())
}

func main() {
	// Creating  Dog and Cat instances.

	dog := Dog{}
	cat := Cat{}

	// Calling the function to print animal sounds.
	PrintAnimalSound(dog)
	PrintAnimalSound(cat)
}
