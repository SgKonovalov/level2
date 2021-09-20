package main

type visitor interface {
	visitForStandartCarVPrice(*standartCarV)
	visitForLuxeCarVPrice(*luxeCarV)
}
