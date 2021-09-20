package definitions

type BodyTypeF interface {
	SetType(name string)
	SetPrice(price int)
	GetType() string
	GetPrice() int
}

type CarF struct {
	Name  string
	Price int
}

func (cf *CarF) SetType(name string) {
	cf.Name = name
}

func (cf *CarF) GetType() string {
	return cf.Name
}

func (cf *CarF) SetPrice(price int) {
	cf.Price = price
}

func (cf *CarF) GetPrice() int {
	return cf.Price
}
