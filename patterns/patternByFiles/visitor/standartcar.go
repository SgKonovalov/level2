package main

type standartCarV struct {
	price int
}

func (scv *standartCarV) accept(v visitor) {
	v.visitForStandartCarVPrice(scv)
}

func (scv *standartCarV) getType() string {
	return "Standart car Price"
}
