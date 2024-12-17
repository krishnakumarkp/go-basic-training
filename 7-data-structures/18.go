package main

import "fmt"

func main() {
	//declaration of map
	var marks map[string]int

	//zero value is nil
	fmt.Println(marks == nil)

	//initialization
	marks = make(map[string]int)

	fmt.Println(marks == nil)

	//or
	marks := make(map[string]int)

	marks["krishna"] = 40
	marks["abishek"] = 25
	marks["saif"] = 30
	//marks["tejaswini"] = 25

	fmt.Println("marks:", marks)
	//size of the map is determined with the len function
	fmt.Println("len:", len(marks))

	value1 := marks["krishna"]
	fmt.Println("marks of krishna: ", value1)

	//element is removed from map with the delte function

	delete(marks, "saif")
	fmt.Println("marks:", marks)

	// //the optional second return value when getting a vlue form a map
	// //indicates if the key was present in the map

	value2, present := marks["tejaswini"]
	fmt.Println("present:", present)

	//fmt.Println("marks of tejaswini:", value2)

	if present {
		fmt.Println("marks of tejaswini:", value2)
	}

	//idiomatic Go style to check if the key exists
	if val, found := marks["tejaswini"]; found {
		fmt.Println("marks of tejaswini:", val)
	}

	//zero value for a non existing key
	value3 := marks["praful"]
	fmt.Println("marks of praful:", value3)

	//map literal
	locations := map[string]string{"krishna": "pune", "abhijit": "gandhinagar"}
	fmt.Println("locations:", locations)

}
