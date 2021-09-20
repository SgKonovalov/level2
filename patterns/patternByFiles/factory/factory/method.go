package factory

import (
	"factory/definitions"
	"fmt"
)

func GetCarF(bodyType string) (definitions.BodyTypeF, error) {
	if bodyType == "SUV" {
		return definitions.NewSUV(), nil
	}
	if bodyType == "Cabrio" {
		return definitions.NewCabrio(), nil
	}
	if bodyType == "Coupe" {
		return definitions.NewCoupe(), nil
	}
	return nil, fmt.Errorf("Wrong body type passed")
}
