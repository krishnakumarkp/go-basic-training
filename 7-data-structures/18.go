package main

import "fmt"

func main() {
	//declaration of map
	var age map[string]int

	//zero value is nil
	fmt.Println(age == nil)

	//initialization
	age = make(map[string]int)

	fmt.Println(age == nil)

	//or
	age := make(map[string]int)

	age["krishna"] = 40
	age["abishek"] = 25
	age["saif"] = 30
	//age["tejaswini"] = 25

	fmt.Println("age:", age)
	//size of the map is determined with the len function
	fmt.Println("len:", len(age))

	value1 := age["krishna"]
	fmt.Println("age of krishna: ", value1)

	//element is removed from map with the delte function

	delete(age, "saif")
	fmt.Println("age:", age)

	// //the optional second return value when getting a vlue form a map
	// //indicates if the key was present in the map

	value2, present := age["tejaswini"]
	fmt.Println("present:", present)

	//fmt.Println("age of tejaswini:", value2)

	if present {
		fmt.Println("age of tejaswini:", value2)
	}

	//idiomatic Go style to check if the key exists
	if val, found := age["tejaswini"]; found {
		fmt.Println("age of tejaswini:", val)
	}

	//zero value for a non existing key
	value3 := age["prafull"]
	fmt.Println("age of prafull:", value3)

	//map literal
	locations := map[string]string{"krishna": "pune", "abhijit": "gandhinagar"}
	fmt.Println("locations:", locations)

}
