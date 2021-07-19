package main

import "fmt"

func main() {
	//iterate over a map
	countries := map[string]string{
		"sk": "Solvakia",
		"ru": "Russia",
		"de": "Germany",
		"no": "Norway",
	}
	for key := range countries {
		fmt.Println(key, "=>", countries[key])
	}

	for key, value := range countries {
		fmt.Printf("countries[%s] = %s\n", key, value)
	}
	// Go map is a reference type. It means that when we assign
	// a reference to a new variable or pass a map to a function,
	// the reference to the map is copied.
	countries2 := countries
	countries2["usa"] = "USA"

	fmt.Println(countries)

}
