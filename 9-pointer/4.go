package main

import "fmt"

func main() {
	//passing array pointer
	var numbers = [3]int{1, 2, 3}
	changeArrayValue(&numbers)

	fmt.Printf("numbers = %v \n", numbers)
}

func changeArrayValue(p *[3]int) {
	(*p)[0] = 4
	(*p)[1] = 5
	(*p)[2] = 6

	// shortcut
	p[0] = 7
	p[1] = 8
	p[2] = 9

}
