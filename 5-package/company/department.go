package company

import "fmt"

func PrintDepartment() {
	fmt.Printf("Engineering department at %s \n", companyName)
}

func init() {
	fmt.Println("Department initialization called")
}
