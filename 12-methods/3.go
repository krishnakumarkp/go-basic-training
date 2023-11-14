package main

import (
	"errors"
	"fmt"
)

// method can be attached to any defined types
//attaching a method to a map

type mark map[string]int

func (m mark) add(k string, v int) error {
	if _, ok := m[k]; ok {
		return errors.New(k + " is already present")
	}
	m[k] = v
	return nil
}

func main() {

	var marksData = make(mark)

	if err := marksData.add("krishna", 90); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("===mark added====")
	}

}
