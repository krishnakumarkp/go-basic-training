// package main

// import "fmt"

// func main() {
// 	var value rune
// 	value = '♬'
// 	fmt.Printf("Character: %c , Decimal: %d,  Unicode: %U", value, value, value)
// }

package main

import "fmt"

func main() {
	// Declare a string containing the word "Hello" in Hindi (नमस्ते)
	str := "नमस्ते"

	// Iterate over each rune in the string
	fmt.Println("Runes in the string:")
	for _, r := range str {
		fmt.Printf("Character: %c , Decimal: %d,  Unicode: %U", r, r, r)
	}
	fmt.Println()

	// Print the length of the string (number of bytes)
	fmt.Println("Length of the string (in bytes):", len(str))

	// Print the length of the string (number of runes)
	fmt.Println("Length of the string (in runes):", len([]rune(str)))
}
