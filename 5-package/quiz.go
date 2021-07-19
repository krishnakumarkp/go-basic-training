package main

import "fmt"

func main() {
	i := 2
	s := "1000"

	if len(s) > 1 {
		i := 4
		i = i + 5
	}
	fmt.Println(i)
}
