package control

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
1) Функция ReadCommandFromConsole() - читает введённую пользователем команду и возвращает её текст.
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
«разбивает» её по ключам и выдаёт bool-значения для каждого ключа, при его наличии.
ВАЖНО: для исключения сбоя в обработке ключей, при помощи strings.TrimLeft(command, "sort"),
убираем слово sort из левой части строки команды.
*/

func ParsingGettingCommand(command string) (kkey, nkey, rkey, ukey, ckey, Mkey, bkey, hkey bool, numOfColumn int) {

	onlyKeys := strings.TrimLeft(command, "sort")

	for _, findKeys := range onlyKeys {

		if findKeys == 'k' {
			kkey = true
		}

		if unicode.IsDigit(findKeys) {
			numOfColumn = int(findKeys - '0')
		}
		if findKeys == 'n' {
			nkey = true
		}
		if findKeys == 'r' {
			rkey = true
		}
		if findKeys == 'u' {
			ukey = true
		}

		if findKeys == 'c' {
			ckey = true
		}
		if findKeys == 'M' {
			Mkey = true
		}
		if findKeys == 'b' {
			bkey = true
		}
		if findKeys == 'h' {
			hkey = true
		}
	}

	return kkey, nkey, rkey, ukey, ckey, Mkey, bkey, hkey, numOfColumn
}

/*
3) Функция  ReadFileNameFromConsole() – запрашивает у пользователя имя файла и выдаёт полученный результат
в виде string.
*/

func ReadFileNameFromConsole() (fileName string) {
	fmt.Println("Insert file name and then press 'Enter', please:")

	myscan := bufio.NewScanner(os.Stdin)
	myscan.Scan()
	fileName = myscan.Text()
	return fileName
}
