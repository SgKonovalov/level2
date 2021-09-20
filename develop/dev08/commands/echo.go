package commands

import (
	"bufio"
	"log"
	"os"
)

/*
Функция ECHO() – считывает текст введённый пользователем и выводит его на консоль, при помощи встроенного
сканера и логов.
*/

func ECHO() {

	log.Println("Insert your text, please")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		usersText := scanner.Text()
		log.Printf("You insered: %s\n", usersText)
	}
}
