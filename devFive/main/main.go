package main

import (
	"fmt"
	"log"

	"exeTwo.devThree/control"
	"exeTwo.devThree/functions"
)

/*
Для работы программы, достаточно вызвать «координатора»:
ChooseOption(ParsingGettingCommand(ReadCommandFromConsole())),
в параметры которому передать ParsingGettingCommand – расшифровка введённых ключей и имени файла.
*/

func main() {

	test, err := functions.ChooseOption(control.ParsingGettingCommand(control.ReadCommandFromConsole()))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You result is:\n%v.", test)

}
