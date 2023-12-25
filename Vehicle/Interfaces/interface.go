package Interfaces

// Vehicle содержит общие поля для всех транспортных средств
type Vehicle struct {
	ID    int
	Brand string
	Model string
	Year  int
}

// AbstractVehicle интерфейс
type AbstractVehicle interface {
	StartEngine()
	StopEngine()
	Accelerate(int)
	Brake()
	DisplayInfo()
}

// Flyable - интерфейс для летающих транспортных средств
type Flyable interface {
	AbstractVehicle
	TakeOff()
	Land()
}

// Swimmable - интерфейс для плавающих транспортных средств
type Swimmable interface {
	AbstractVehicle
	StartSwimming()
	StopSwimming()
}
