package reporting

import (
	"fmt"
	"krishnakumarkp/srp/solution/internal/employee"
)

type Reporter struct{}

func (r Reporter) Report(emp employee.Employee) {
	fmt.Printf("%s worked %.2f hours\n", emp.Name, emp.Hours)
}
