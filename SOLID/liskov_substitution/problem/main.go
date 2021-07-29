package main

import (
	"fmt"
	"liskovsubstitution/petshop/pets"
)

//Liskov Substitution Principle
//Subtypes must be substitutable for their base types

// in mathematics, a Square is a Rectangle. Indeed it is a specialization of a rectangle.
// The "is a" makes you want to model this with inheritance. However if in code you made Square derive from Rectangle,
// then a Square should be usable anywhere you expect a Rectangle.
// imagine you had SetWidth and SetHeight methods on your Rectangle base class; this seems perfectly logical.
// However if your Rectangle reference pointed to a Square, then SetWidth and SetHeight doesn't make sense

// public class Bird{
//     public void fly(){}
// }
// public class Duck extends Bird{}
// The duck can fly because it is a bird, but what about this:
// public class Ostrich extends Bird{}
// Ostrich is a bird, but it can't fly

// Good exmaple
// public class Bird{}
// public class FlyingBirds extends Bird{
//     public void fly(){}
// }
// public class Duck extends FlyingBirds{}
// public class Ostrich extends Bird{}

// In a class based language, Liskov’s substitution principle is commonly interpreted as a specification
// for an abstract base class with various concrete subtypes. But Go does not have classes, or inheritance,
// so substitution cannot be implemented in terms of an abstract class hierarchy.

//Instead, substitution is the purview of Go’s interfaces. in go interfaces are satisfied implicitly,
// rather than explicitly, and this has a profound impact on how they are used within the language.

// Go does not have inheritance insted it uses composition.
//composition does not allow to substitute parent struct by child
//So this code will panic :

func main() {

	pet := pets.Dog{}
	//pet := pets.Pet{}
	pet.Name = "Balto"
	WelcomeGuest(pet)

}

func WelcomeGuest(p pets.Pet) {
	fmt.Println(p.Greet())
}
