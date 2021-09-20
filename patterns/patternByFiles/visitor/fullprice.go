package main

import "fmt"

type fullPrice struct{}

func (fp *fullPrice) visitForStandartCarVPrice(scv *standartCarV) {
	fmt.Printf("Standart Car is %d\n", scv.price+500)
}

func (fp *fullPrice) visitForLuxeCarVPrice(lcv *luxeCarV) {
	fmt.Printf("Luxe Car is %d\n", lcv.price+1500)
}
