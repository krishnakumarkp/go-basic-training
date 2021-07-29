package main

import "interfacesegregation/vehiclefactory/vehicles"

//Interface Segregation Principle
//Clients should not be forced to depend on methods they do not use.

func main() {
	//we separated the 2 different behaviors as components CDPlayer and Accelerator.defined as 2 interfaces,
	//and implimented by 2 different struct.

	vehicle := vehicles.Vehicle{}
	vehicle.Accelerate()
	vehicle.PlayCD()

	bus := vehicles.Bus{} // compose with VehicleAccelerator
	bus.Accelerate()
	//bus.PlayCD()
	//Bus does not have PlayCD now.
	//so here  bus is now not forced to depend on methods they do not use. so this follows Interface Segregation Principle

	// The Interface Substitution Principle encourages you to define functions and methods
	// that depend only on the behaviour that they need. If your function only requires a parameter of an interface type
	// with a single method, then it is more likely that this function has only one responsibility.

}
