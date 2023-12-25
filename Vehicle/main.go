package main

import (
	"Vehicle/Aircrafts"
	"Vehicle/Boats"
	"Vehicle/Cars"
	"fmt"
)

func main() {
	car := Cars.NewCar(1, "Toyota", "Camry", 2022, 60, 90)
	aircraft := Aircrafts.NewAircraft(2, "Boeing", "747", 2010, 10000, false)
	boat := Boats.NewBoat(3, "Yamaha", "242X", 2021, 40, false)

	car.DisplayInfo()
	car.StartEngine()
	car.Accelerate(80)
	car.Brake()
	car.Refuel(20)
	car.Refuel(20)
	car.DisplayInfo()

	fmt.Println()

	aircraft.DisplayInfo()
	aircraft.StartEngine()
	aircraft.TakeOff(1000)
	aircraft.TakeOff(2000)
	aircraft.Land()
	aircraft.StopEngine()
	aircraft.DisplayInfo()

	fmt.Println()

	boat.DisplayInfo()
	boat.StartEngine()
	boat.StartSwimming(80)
	boat.StopSwimming()
	boat.StopSwimming()
	boat.StopEngine()
	boat.DisplayInfo()
}
