package difinitions

type StandartCar struct {
	StearingWheel  string
	Seats          string
	Audio          string
	AirConditioner string
}

func NewStandart() *StandartCar {
	return &StandartCar{}
}

func (sc *StandartCar) InstallStearingWheel() {
	sc.StearingWheel = "Polyurethane stering wheel"
}

func (sc *StandartCar) InstallSeats() {
	sc.Seats = "Cloth seats"
}

func (sc *StandartCar) InstallAudio() {
	sc.Audio = "No audio"
}

func (sc *StandartCar) InstallAirConditioner() {
	sc.AirConditioner = "No air conditioner"
}

func (sc *StandartCar) GetCar() Car {
	return Car{
		StearingWheel:  sc.StearingWheel,
		Seats:          sc.Seats,
		Audio:          sc.Audio,
		AirConditioner: sc.AirConditioner,
	}
}

type LuxeCar struct {
	StearingWheel  string
	Seats          string
	Audio          string
	AirConditioner string
}

func NewLuxeCarBuilder() *LuxeCar {
	return &LuxeCar{}
}

func (lc *LuxeCar) InstallStearingWheel() {
	lc.StearingWheel = "Leather stering wheel"
}

func (lc *LuxeCar) InstallSeats() {
	lc.Seats = "Leather seats"
}

func (lc *LuxeCar) InstallAudio() {
	lc.Audio = "Luxe audio"
}

func (lc *LuxeCar) InstallAirConditioner() {
	lc.AirConditioner = "Climate-control"
}

func (lc *LuxeCar) GetCar() Car {
	return Car{
		StearingWheel:  lc.StearingWheel,
		Seats:          lc.Seats,
		Audio:          lc.Audio,
		AirConditioner: lc.AirConditioner,
	}
}
