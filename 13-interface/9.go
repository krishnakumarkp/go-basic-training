package main

import "fmt"

type Person struct {
	name string
}

func main() {
	var num int8 = 1
	describe(num)

	var str string = "krishna"
	describe(str)

	per := Person{
		name: "Naveen",
	}

	describe(per)

}

func describe(i interface{}) {
	fmt.Println(i)
	fmt.Printf("Interface had type %T value %v\n", i, i)
}
