package main

import "fmt"

type Worker interface {
	Work()
}

func main() {
	var w Worker

	if w == nil {
		fmt.Printf("w is nil and has type %T value %v\n", w, w)
	}

}
