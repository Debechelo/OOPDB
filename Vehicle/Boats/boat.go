package Boats

import (
	"Vehicle/Interfaces"
	"fmt"
)

type Boat struct {
	Interfaces.Vehicle
	MaxSpeed  int
	IsSailing bool
}

func NewBoat(id int, brand, model string, year, maxSpeed int, isSailing bool) *Boat {
	if year < 1900 {
		fmt.Println("Wrong year")
	}
	if maxSpeed < 0 {
		fmt.Println("Wrong speed")
	}
	return &Boat{Vehicle: Interfaces.Vehicle{ID: id, Brand: brand, Model: model, Year: year},
		MaxSpeed:  maxSpeed,
		IsSailing: isSailing,
	}
}

func (b *Boat) StartEngine() {
	fmt.Println("Boat engine started.")
}

func (b *Boat) StopEngine() {
	fmt.Println("Boat engine stopped.")
}

func (b *Boat) StartSwimming(speed int) {
	if speed < 0 {
		fmt.Println("Boat won't go.")
	} else {
		b.Accelerate(speed)
		b.IsSailing = true
	}
}

func (b *Boat) Accelerate(speed int) {
	fmt.Printf("Boat started swimming to %d km/h.\n", speed)
}

func (b *Boat) StopSwimming() {
	if !b.IsSailing {
		fmt.Println("Boat is not swimming.")
		return
	}
	b.Brake()
	b.IsSailing = false
}

func (b *Boat) Brake() {
	fmt.Println("Boat stopped swimming.")
}

func (b *Boat) DisplayInfo() {
	sailingStatus := "at anchor"
	if b.IsSailing {
		sailingStatus = "sailing"
	}
	fmt.Printf("Boat: %d %s %s (%d), Max Speed: %d knots, Status: %s\n",
		b.ID,
		b.Brand,
		b.Model,
		b.Year,
		b.MaxSpeed,
		sailingStatus)
}
