package components

import "fmt"

type VehicleAccelerator struct {
}

func (v VehicleAccelerator) Accelerate() {
	fmt.Println("Accelerating ....")
}
