package main

import (
	"fmt"
)

func main() {
	// // //An array can be partially assigned
	// // var a [5]int = [5]int{10, 20, 30}
	// // fmt.Println(a)

	// // //In the array literal, we can provide the index of the element.
	// // b := [5]int{1: 6, 2: 7, 4: 9}
	// // fmt.Println(b)

	// // //go can infer array length, for this , we use the ellipses ... operator
	// // c := [...]int{1, 2, 3, 4, 5, 6}
	// // fmt.Println(c)

	// //colon character to retrieve a portion of the array
	words := [...]string{"falcon", "sky", "earth", "cloud", "fox"}

	// //index o to 2 not including 2
	// fmt.Println("0:2", words[0:2])
	// //index 1 to 4 not including 4
	// fmt.Println("1:4", words[1:4])
	// //index 1 to end of array
	// fmt.Println("1:", words[1:])
	// //from beginning of array to index 2 not including 2
	// fmt.Println(":2", words[:2])

	//Different ways of Iterating through array
	// for i := 0; i < len(words); i++ {
	// 	fmt.Println(words[i])
	// }

	for index, value := range words {
		fmt.Println(index, "=>", value)
	}

	// j := 0

	// for range words {
	// 	fmt.Println(words[j])
	// 	j++
	// }

	//Multidimensional array in Golang

	// create a 2 dimensional array
	arrayInteger := [2][2]int{{1, 2}, {3, 4}}

	// access the values of 2d array
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arrayInteger[i][j])
		}
	}

}
