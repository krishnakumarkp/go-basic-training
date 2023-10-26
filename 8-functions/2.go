package main

import (
	"errors"
	"fmt"
)

//Error handling in Go

func main() {

	result := divide(3, 0)
	fmt.Printf("result : %v", result)

}

func divide(x int, y int) int {
	return x / y
}

//how will we handle this error
// other programming languages provides the conventional try/catch method 

try {
	result := divide(3, 0);
}
catch(Exception $e) {
	echo 'Message: ' .$e->getMessage();
}

//Go does not provide conventional try/catch method to handle the errors,
//instead, errors are returned as a normal return value



func main() {
	//error handling
	result, err := divide(5, 0)

	if err != nil {
		fmt.Printf("error occured : %v", err)
	} else {
		fmt.Printf("result : %v", result)
	}
	//_, err = divide(5, 0)

	//Idiomatic Go
	//error handling without else by return earnly
	if err != nil {
		fmt.Printf("error occured : %v", err)
		return
	}
	fmt.Printf("result : %v", result)

	//blank identifier

}

func divide(a, b float64) (float64, error) {
	var result float64
	if b == 0 {
		return result, errors.New("can't divide by 0")
	}
	result = a / b
	return result, nil
}

//probelem with this approach no Stack Traces
//Without stack traces, it is very difficult to debug when you have no idea where the error occurs. 
//The error could pass through 10 function calls before it gets to the code that prints it out.


func divideNumbers(x, y int) (int, error) {
	return divide(x, y)
}


func main() {

	result, err := divideNumbers(3, 0)

	if err != nil {
		fmt.Printf("error occured : %v", err)
	} else {
		fmt.Printf("result : %v", result)
	}
}

//Many people have this exact same problem and they created awesome projects to handle this issue.
// palantir/stacktrace, go-erros/errors, and pkg/errors are one of them.

//let us change the import statement from “errors” to “github.com/pkg/errors” 

go get github.com/pkg/errors


import (
	"github.com/pkg/errors"
	"fmt"
 )

 //If you want to print stack traces instead of a plain error message, you have to use %+v instead of %v in the format pattern, 

