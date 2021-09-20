package main

import "exeTwo.devTen/command"

/*
Для работы программы, достаточно вызвать:
command.ChooseOption(command.ParsingGettingCommand(command.ReadCommandFromConsole()))
*/

func main() {

	command.ChooseOption(command.ParsingGettingCommand(command.ReadCommandFromConsole()))
}
