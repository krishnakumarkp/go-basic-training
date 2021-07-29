package main

import "interfacesegregation/vehiclefactory/vehicles"

//Interface Segregation Principle
//Clients should not be forced to depend on methods they do not use.

func main() {
	vehicle := vehicles.Vehicle{}
	vehicle.Accelerate()
	vehicle.PlayCD()

	bus := vehicles.Bus{} //extended our vehicle to implement a Bus
	bus.Accelerate()
	//bus.PlayCD()
	//Bus does not need to PlayCD. But he will be obliged to have the method, even if it will not be used.
	//so here  bus forced to depend on methods they do not use. so it breaks Interface Segregation Principle

}
