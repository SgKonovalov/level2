package control

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
1) Функция ReadCommandFromConsole() - читает введённую пользователем команду и возвращает её текст;
*/

func ReadCommandFromConsole() (command string) {

	fmt.Println("Insert your command and then press 'Enter', please:")

	myscan := bufio.NewScanner(os.Stdin)
	myscan.Scan()
	command = myscan.Text()
	return command
}

/*
2) Функция ParsingGettingCommand() – принимает от ReadCommandFromConsole текст команды,
«разбивает» её по ключам и выдаёт bool-значения для каждого ключа, при его наличии;
*/

func ParsingGettingCommand(command string) (Akey, Bkey, Ckey, ckey, ikey, vkey, Fkey, nkey bool, numOfStrings int) {

	onlyKeys := strings.TrimLeft(command, "grep")

	for _, findKeys := range onlyKeys {

		if findKeys == 'A' {
			Akey = true
		}

		if unicode.IsDigit(findKeys) {
			numOfStrings = int(findKeys - '0')
		}

		if findKeys == 'B' {
			Bkey = true
		}
		if findKeys == 'C' {
			Ckey = true
		}
		if findKeys == 'c' {
			ckey = true
		}

		if findKeys == 'i' {
			ikey = true
		}
		if findKeys == 'v' {
			vkey = true
		}

		if findKeys == 'F' {
			Fkey = true
		}
		if findKeys == 'n' {
			nkey = true
		}
	}

	return Akey, Bkey, Ckey, ckey, ikey, vkey, Fkey, nkey, numOfStrings
}

/*
3) Функция  ReadFileNameFromConsole() – запрашивает у пользователя имя файла и выдаёт полученный результат
в виде string;
*/

func ReadTextFromConsole() (searchingText string) {
	fmt.Println("Insert searching text or part and then press 'Enter', please:")

	myscan := bufio.NewScanner(os.Stdin)
	myscan.Scan()
	searchingText = myscan.Text()
	return searchingText
}

/*
4) Функция ReadTextFromConsole() –запрашивает у пользователя фрагмент текста, который он хочет найти.
*/

func ReadFileNameFromConsole() (fileName string) {
	fmt.Println("Insert file name and then press 'Enter', please:")

	myscan := bufio.NewScanner(os.Stdin)
	myscan.Scan()
	fileName = myscan.Text()
	return fileName
}
