package main

import (
	"fmt"
	"go-training/package/company"
	"go-training/package/home"
)

func main() {
	company.PrintCompany()
	company.PrintDepartment()
	home.PrintHomeName()
}

func init() {
	fmt.Println("Main initialization called")
}
