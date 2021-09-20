package definition

import (
	"strategy/action"
)

type Car struct {
	CarC action.CarColour
}

func (c *Car) Painting() {
	c.CarC.Coloring()
}

func (cmc *Car) SetState(carPSt action.CarColour) {
	cmc.CarC = carPSt
}
