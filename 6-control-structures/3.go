package main

import "fmt"

func main() {
	sum := 1
	//The init and post statements are optional.Which makes it while
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
