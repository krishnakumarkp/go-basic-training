package company

import (
	"fmt"
)

const GrowthPercentage = 6.1415926535897932384

var companyName string = "cybage"

func PrintCompany() {
	fmt.Println(companyName)
}

func init() {
	fmt.Println("Company initialization called")
}
