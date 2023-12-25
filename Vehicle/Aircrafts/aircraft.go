package Aircrafts

import (
	"Vehicle/Interfaces"
	"fmt"
)

type Aircraft struct {
	Interfaces.Vehicle
	maxAltitude int
	isFlying    bool
}

func NewAircraft(id int, brand, model string, year, maxAltitude int, isFlying bool) *Aircraft {
	if year < 1900 {
		fmt.Println("Wrong year")
	}
	if maxAltitude < 0 {
		fmt.Println("Wrong altitude")
	}
	return &Aircraft{Vehicle: Interfaces.Vehicle{ID: id, Brand: brand, Model: model, Year: year},
		maxAltitude: maxAltitude,
		isFlying:    isFlying,
	}
}

func (a *Aircraft) StartEngine() {
	fmt.Println("Aircraft engine started.")
}

func (a *Aircraft) StopEngine() {
	fmt.Println("Aircraft engine stopped.")
}

func (a *Aircraft) TakeOff(altitude int) {
	if altitude > a.maxAltitude || altitude < 0 {
		fmt.Println("Aircraft won't take off.")
	} else {
		a.Accelerate(altitude)
		a.isFlying = true
	}
}

func (a *Aircraft) Accelerate(altitude int) {
	fmt.Printf("Aircraft took off on altitude: %d \n", altitude)
}

func (a *Aircraft) Land() {
	if !a.isFlying {
		fmt.Println("Aircraft is not flying.")
		return
	}
	a.Brake()
	a.isFlying = false
}

func (a *Aircraft) Brake() {
	fmt.Println("Aircraft landed.")
}

func (a *Aircraft) DisplayInfo() {
	flyingStatus := "on the ground"
	if a.isFlying {
		flyingStatus = "in the air"
	}
	fmt.Printf("Aircraft: %s %s %s (%d), Altitude: %d meters, Status: %s\n",
		a.ID,
		a.Brand,
		a.Model,
		a.Year,
		a.maxAltitude,
		flyingStatus)
}
