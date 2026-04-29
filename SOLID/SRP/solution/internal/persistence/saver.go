package persistence

import (
	"fmt"
	"krishnakumarkp/srp/solution/internal/employee"
)

type Saver struct{}

func (s Saver) Save(emp employee.Employee) {
	fmt.Printf("Saving %s to database\n", emp.Name)
}
