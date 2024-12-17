package main

import (
	"fmt"
	"go-training/alias-use/users"
)

func main() {
	data := make(users.Userdata)

	users.AddUser(data, 1, "krishna")

	fmt.Println(users.GetUser(data, 1))

	//now new users of the package can use Userlist

	data2 := make(users.Userlist)

	users.AddUser(data2, 1, "krishna")

	fmt.Println(users.GetUser(data2, 1))
}
