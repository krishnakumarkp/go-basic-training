package components

import "fmt"

type VehicleCDPlayer struct {
}

func (v VehicleCDPlayer) PlayCD() {
	fmt.Println("Playnig Guns n Roses")
}
