package main

import "fmt"

// Anonymous fields
// You can define a struct type without declaring any field names.
// You have to just define the field data types and Go will use the data type declarations
// as the field names

type Data struct {
	string
	int
	bool
}

func main() {
	sample1 := Data{"Monday", 2021, true}
	fmt.Printf("sample1 = %+v\n", sample1)

}
