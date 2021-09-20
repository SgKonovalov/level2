package context

import (
	"state/state"
)

type CarMovingContext struct {
	Cs state.CarState
}

func (cmc *CarMovingContext) SetState(cs state.CarState) {
	cmc.Cs = cs
}

func (cmc CarMovingContext) Action() {
	cmc.Cs.Action()
}
