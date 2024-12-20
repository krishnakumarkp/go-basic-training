package main

import (
	"database/sql"
	"fmt"
)

func main() {
	//A panic is smimiler to an exception which can occur at the runtime.
	//a panic can be thrown by the runtime or can be deliberately thrown
	//by the programmer.

	a := []int{1, 2, 3}
	//fmt.Println(accessElement(a, 3))
	fmt.Println(accessElement2(a, 3))
}

// thrown by the runtime
func accessElement(a []int, index int) int {
	return a[index]
}

// deliberately thrown
func accessElement2(a []int, index int) int {
	if len(a) > index {
		return a[index]
	} else {
		panic("Out of bound condition")
	}
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	return db
}
