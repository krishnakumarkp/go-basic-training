package main

import "fmt"

// "Slice" operator with the syntax slice[low:high]
func main() {
	words := []string{"falcon", "sky", "earth", "cloud", "fox"}

	//no step value
	fmt.Println(words[0:2])
	fmt.Println(words[1:4])
	fmt.Println(words[1:])
	fmt.Println(words[:3])
	fmt.Println(words[:len(words)-1])

	//Different ways of Iterating through slice
	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])
	}

	for index := range words {
		fmt.Println(words[index])
	}

	for index, value := range words {
		fmt.Println(index, "=>", value)
	}

	j := 0

	for range words {
		fmt.Println(words[j])
		j++
	}

}
