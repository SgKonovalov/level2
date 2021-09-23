package commands

import (
	"os/exec"
)

/*
Функция EXEC() – получает путь к файлу, который хочет открыть пользователь и при помощи
функции exec.Command и cmd.Run() выполняет запрос на запуск программы.
*/

func EXEC(pathToExeFile string) (err error) {

	cmd := exec.Command(pathToExeFile)
	cmd.Run()

	return nil
}
