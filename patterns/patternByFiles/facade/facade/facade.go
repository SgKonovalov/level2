package facade

import (
	"strings"

	"facade/api"
)

type forCar interface {
	Actions() string
}

func NewCar() *Car {
	return &Car{
		stearingWheel:  &api.StearingWheel{},
		seats:          &api.Seats{},
		audio:          &api.Audio{},
		airConditioner: &api.AirConditioner{},
	}
}

type Car struct {
	stearingWheel  *api.StearingWheel
	seats          *api.Seats
	audio          *api.Audio
	airConditioner *api.AirConditioner
}

func (c *Car) Actions() string {
	result := []string{
		c.stearingWheel.TurnWheels(),
		c.seats.RaiseBack(),
		c.audio.TurnOn(),
		c.airConditioner.TurnOffAC(),
	}
	return strings.Join(result, "\n")
}
