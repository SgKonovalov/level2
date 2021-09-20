package api

type StearingWheel struct {
}

func (sw *StearingWheel) TurnWheels() string {
	return "Turns wheels at write"
}

type Seats struct {
}

func (st *Seats) RaiseBack() string {
	return "Raise the back of seat"
}

type Audio struct {
}

func (a *Audio) TurnOn() string {
	return "Turn on audio"
}

type AirConditioner struct{}

func (hl *AirConditioner) TurnOffAC() string {
	return "Turn off A/C"
}
