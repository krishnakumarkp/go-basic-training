package vehicles

import (
	"interfacesegregation/vehiclefactory/components"
)

type Vehicle struct {
	components.VehicleAccelerator
	components.VehicleCDPlayer
}
