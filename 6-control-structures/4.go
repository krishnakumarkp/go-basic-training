package main

import "fmt"

func main() {
	sum := 1
	//condition is also optional which makes it for ever
	for {
		sum += sum
		//skip odd numbers
		if sum%2 != 0 {
			continue
		}
		fmt.Println(sum)
		if sum > 10000 {
			break
		}
	}
}
