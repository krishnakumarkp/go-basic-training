package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (u *Person) IncrementAge() {
	u.Age++
}

//problem with iterating slice of objects using range

func main() {

	users := []Person{
		{"Abhinav", 25},
		{"Priyanka", 27},
		{"Darshan", 27},
	}

	//you get a copy of the slice value in user.
	//So if you change it you are changing this copy

	for _, user := range users {
		user.IncrementAge()
	}

	//this can be written in different way

	for i := range users {
		users[i].IncrementAge()
	}

	fmt.Println("Employee age after change")

	fmt.Printf("%+v\n", users)
}
