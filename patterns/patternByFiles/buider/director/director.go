package director

import (
	"builder/builder"
	"builder/difinitions"
)

type Director struct {
	carBuilder builder.CarBuilder
}

func NewDirector(cb builder.CarBuilder) *Director {
	return &Director{
		carBuilder: cb,
	}
}

func (d *Director) SetCarBuilder(cb builder.CarBuilder) {
	d.carBuilder = cb
}

func (d *Director) ProduceCar() difinitions.Car {
	d.carBuilder.InstallStearingWheel()
	d.carBuilder.InstallSeats()
	d.carBuilder.InstallAudio()
	d.carBuilder.InstallAudio()
	d.carBuilder.InstallAirConditioner()
	return d.carBuilder.GetCar()
}
