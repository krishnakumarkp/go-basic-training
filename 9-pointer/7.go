package main

import "fmt"

func main() {
	// passing map by reference
	// maps are actually reference types and therefore
	// there is a very limited use-case of using pointers.
	original := make(map[string]int)

	original["krishna"] = 40
	original["abhijit"] = 30

	//modify(&original)
	modify2(original)

	fmt.Println(original)
}

func modify(m *map[string]int) {
	(*m)["krishna"] = 20
}

func modify2(m map[string]int) {
	m["krishna"] = 30
	m["saif"] = 25
}

// maps are actually reference types and therefore
// there is a very limited use-case of using pointers.
