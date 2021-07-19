package main

import (
	"fmt"
)

// Type switches
// A type switch performs several type assertions in series and runs the first
// case with a matching type.

type Programmer interface {
	Code()
}

type Person struct {
	name string
}

func (p Person) Code() {
	fmt.Println("I can code in go")
}

func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("i am string and my value is %s \n", i.(string))
	case int:
		fmt.Printf("i am int and my value is %d \n", i.(int))
	case Person:
		fmt.Printf("i am person and my value is %+v \n", i.(Person))
	case Programmer:
		fmt.Printf("i am programmer and my value is %+v \n", i.(Person))

	default:
		fmt.Println("i stored something else", i)
	}
}

func main() {
	var p Programmer

	p = Person{"krishna"}

	explain("Hello World")
	explain(52)
	explain(true)
	explain(p)
}
