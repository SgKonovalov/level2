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

func ParsingGettingCommand(command string) (fkey, dkey, skey bool, numOfColumn int) {

	onlyKeys := strings.TrimLeft(command, "cut")

	for _, findKeys := range onlyKeys {

		if unicode.IsDigit(findKeys) {
			numOfColumn = int(findKeys - '0')
		}

		if findKeys == 'f' {
			fkey = true
		}

		if findKeys == 'd' {
			dkey = true
		}

		if findKeys == 's' {
			skey = true
		}

	}

	return fkey, dkey, skey, numOfColumn
}

/*
3) Функция  ReadFileNameFromConsole() – запрашивает у пользователя имя файла и выдаёт полученный результат
в виде string;
*/

func ReadFileNameFromConsole() (fileName string) {
	fmt.Println("Insert file name and then press 'Enter', please:")

	myscan := bufio.NewScanner(os.Stdin)
	myscan.Scan()
	fileName = myscan.Text()
	return fileName
}

/*
4) Функция ReadNewSeparator() –запрашивает у пользователя новый разделитель для применения с ключом --d.
*/

func ReadNewSeparator() (separator string) {
	fmt.Println("Insert new separator and then press 'Enter', please:")

	myscan := bufio.NewScanner(os.Stdin)
	myscan.Scan()
	separator = myscan.Text()
	if len(separator) > 1 {
		separator = "Too much arguments fo separator! Expected only one, like ':'."
		return separator
	}

	return separator
}
