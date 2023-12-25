package Cars

import (
	"Vehicle/Interfaces"
	"fmt"
)

type Car struct {
	Interfaces.Vehicle
	FuelCapacity     int
	CurrentFuelLevel int
}

func NewCar(id int, brand, model string, year, currentFuelLevel, fuelCapacity int) *Car {
	if year < 1900 {
		fmt.Println("Wrong year")
	}
	if currentFuelLevel < 0 {
		fmt.Println("Wrong fuel level")
	}
	if fuelCapacity < 0 {
		fmt.Println("Wrong fuel capacity")
	}
	return &Car{Vehicle: Interfaces.Vehicle{ID: id, Brand: brand, Model: model, Year: year},
		FuelCapacity:     fuelCapacity,
		CurrentFuelLevel: currentFuelLevel,
	}
}

func (c *Car) StartEngine() {
	fmt.Println("Car engine started.")
}

func (c *Car) StopEngine() {
	fmt.Println("Car engine stopped.")
}

func (c *Car) Accelerate(speed int) {
	if speed < 0 {
		fmt.Println("Car won't go.")
	} else {
		fmt.Printf("Car accelerated to %d km/h.\n", speed)
	}
}

func (c *Car) Brake() {
	fmt.Println("Car stopped.")
}

func (c *Car) Refuel(liters int) {
	if liters < 0 || c.CurrentFuelLevel+liters > c.FuelCapacity {
		fmt.Println("Wrong refuel")
		return
	}
	c.CurrentFuelLevel += liters
	fmt.Printf("Car refueled with %d liters.\n", liters)
}

func (c *Car) DisplayInfo() {
	fmt.Printf("Car: %d %s %s (%d), Fuel: %d/%d liters\n",
		c.ID, c.Brand,
		c.Model,
		c.Year,
		c.CurrentFuelLevel,
		c.FuelCapacity)
}
