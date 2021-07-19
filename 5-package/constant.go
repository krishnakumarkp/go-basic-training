package main

import (
	"fmt"
	"go-training/package/company"
)

const b = 10

func main() {
	company.PrintCompany()
	company.PrintDepartment()
	//typed and untyped constant
	const a int64 = 8
	//local constant with same name as global constant
	const b = 8

	var c int64
	c = a
	c = b
	fmt.Println(c)

	//problem with typed constant
	//var growth float64
	var growth float32
	//var growth float16
	growth = company.GrowthPercentage
	fmt.Println("growth percentage = ", growth)
}
