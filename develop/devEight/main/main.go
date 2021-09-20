package main

import "fmt"

/*
Для работы программы достаточно вызвать chooseCommand(ParsingGettingCommand(ReadCommandFromConsole()))
*/

func main() {
	for {
		result, err := chooseCommand(ParsingGettingCommand(ReadCommandFromConsole()))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(result)
	}
}
