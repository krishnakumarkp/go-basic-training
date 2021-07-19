package main

import (
	"fmt"
	"gotraining/struct/company"
)

//we can also control which fields of an exported struct are visible outside the
// package (or exported). To export the field names of a struct, we’ve to follow
// the same uppercase letter approach.

func main() {
	emp1 := company.Employee{
		FirstName: "Krishna",
		LastName:  "kumar",
		Salary:    1234,
	}
	fmt.Printf("emp1 = %+v\n", emp1)
}
