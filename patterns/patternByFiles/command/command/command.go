package command

import (
	"command/definition"
)

type Command interface {
	Execute()
}

type StartEngineC struct {
	EngineC definition.EngineC
}

func (sec *StartEngineC) Execute() {
	sec.EngineC.StartEngineC()
}

type StopEngineC struct {
	EngineC definition.EngineC
}

func (stec *StopEngineC) Execute() {
	stec.EngineC.StopEngineC()
}
