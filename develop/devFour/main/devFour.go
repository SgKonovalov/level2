package main

import (
	"log"

	base "exeTwo.devFour/base"
)

var sourceOne = []string{"окно", "Тяпка", "листок"}

/*var sourceTwo = []string{"окно", "Тяпка", "Листок"}
var sourceThree = []string{"окно", "Тяпка", "щиток"}
var sourceFour = []string{"окно", "тряпка", "щиток"}
var sourceFive = []string{"столик", "пятка"}*/

func main() {

	test, err := base.GetTheResult(base.Matcher(sourceOne))
	if err != nil {
		log.Fatal(err)
	}

	base.PrintResult(test)
}
