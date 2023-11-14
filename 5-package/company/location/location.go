package location

import (
	"fmt"
	//"go-training/package/company"
)

var locationName string = "pune"

func PrintLocation() {
	fmt.Println(locationName)
	//company.PrintCompany()
}

func init() {
	fmt.Println("location initialization called")
}
