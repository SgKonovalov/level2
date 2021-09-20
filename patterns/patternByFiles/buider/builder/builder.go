package builder

import (
	"builder/difinitions"
)

type CarBuilder interface {
	InstallStearingWheel()
	InstallSeats()
	InstallAudio()
	InstallAirConditioner()
	GetCar() difinitions.Car
}

func GetCarBuilder(builderType string) CarBuilder {
	if builderType == "standart" {
		return &difinitions.StandartCar{}
	}

	if builderType == "luxe" {
		return &difinitions.LuxeCar{}
	}
	return nil
}
