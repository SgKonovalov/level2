package definitions

type SUV struct {
	CarF
}

func NewSUV() BodyTypeF {
	return &SUV{
		CarF: CarF{
			Name:  "SUV",
			Price: 100000,
		},
	}
}

type Cabrio struct {
	CarF
}

func NewCabrio() BodyTypeF {
	return &SUV{
		CarF: CarF{
			Name:  "Cabrio",
			Price: 150000,
		},
	}
}

type Coupe struct {
	CarF
}

func NewCoupe() BodyTypeF {
	return &SUV{
		CarF: CarF{
			Name:  "Coupe",
			Price: 200000,
		},
	}
}
