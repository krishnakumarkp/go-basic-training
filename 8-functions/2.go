package main

import (
	"errors"
	"fmt"
)

//Error handling in Go
//Go does not provide conventional try/catch method to handle the errors,
//instead, errors are returned as a normal return value

// try {
// 	checkNum(2);
// }
// catch(Exception $e) {
// 	echo 'Message: ' .$e->getMessage();
// }

func main() {
	//error handling
	result, err := divide(5, 0)

	if err != nil {
		fmt.Println("error occured : ", err)
	} else {
		fmt.Println("result : ", result)
	}
	//_, err = divide(5, 0)

	//Idiomatic Go
	//error handling without else by return earnly
	if err != nil {
		fmt.Println("error occured : ", err)
		return
	}
	fmt.Println("result : ", result)

	//balnk identifier

}

func divide(a, b float64) (float64, error) {
	var result float64
	if b == 0 {
		return result, errors.New("can't divide by 0")
	}
	result = a / b
	return result, nil
}
