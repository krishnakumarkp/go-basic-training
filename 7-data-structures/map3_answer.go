package main

import (
	"fmt"
	"sort"
)

func main() {
	// Initialize map with make function
	countryCodes := make(map[string]string)

	// Add data as key/value pairs
	countryCodes["Norway"] = "NOR"
	countryCodes["Thailand"] = "THA"
	countryCodes["Zimbabwe"] = "ZWE"
	countryCodes["United States"] = "USA"
	countryCodes["India"] = "IND"

	fmt.Println("Before sorting")
	//iterate thrugh map and print keyy and values
	for k, v := range countryCodes {
		fmt.Println(k, " => ", v)
	}

	// Slice for specifying the order of the map
	var keys []string
	// Appending keys of the map
	for k := range countryCodes {
		keys = append(keys, k)
	}
	//sort.Strings sorts a slice of strings.
	sort.Strings(keys)
	fmt.Println("After sorting")
	// Iterate over the slice and print map in order
	for _, k := range keys {
		fmt.Println(k, " => ", countryCodes[k])
	}
}
