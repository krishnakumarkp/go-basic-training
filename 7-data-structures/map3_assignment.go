package main

import (
	"fmt"
	"sort"
)

//assignment
//Create a map of country and country codes and print it
//sorted in alphabetical order of the key

func main() {
	// Initialize map with make function
	// your code here

	// Add data as key/value pairs
	countryCodes["Norway"] = "NOR"
	countryCodes["Thailand"] = "THA"
	countryCodes["Zimbabwe"] = "ZWE"
	countryCodes["United States"] = "USA"
	countryCodes["India"] = "IND"

	fmt.Println("Before sorting")
	//iterate thrugh map and print key and values
	// your code here

	// create a slice named keys for storing the keys of the map
	// your code here

	// Append keys of the map to key slice
	// your code here

	//sort.Strings sorts a slice of strings.
	sort.Strings(keys)
	fmt.Println("After sorting")
	// Iterate over the slice using for loop and print map in order
	// your code here
}
