package vehicles

import "fmt"

type Vehicle struct {
}

func (v Vehicle) Accelerate() {
	fmt.Println("Accelerating ....")
}
func (v Vehicle) PlayCD() {
	fmt.Println("Playnig Guns n Roses")
}
