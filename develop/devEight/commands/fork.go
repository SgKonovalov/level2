package commands

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

/*
Функция FORK() – запускает копию текущего процесса:
1) При помощи os.Getwd() – находим путь к текущей директории;
2) При помощи ioutil.ReadDir() – перебираем все файлы в текущей директории, передавая в качестве аргумента,
путь из п. 1;
3) Перебираем файлы в текущей директории. Если находим файл с расширением .exe,
присваиваем его имя переменной execFile;
4) При помощи fmt.Sprintf() – конструируем строку по формуле:
«путь к файлу» + «\» + «имя файла с расширением .exe».
5) При помощи функций exec.Command и cmd.Run() – запускаем .exe.
*/

func FORK() (err error) {

	var execFile string

	path, err := os.Getwd()
	if err != nil {
		return errors.New(cantFoundCurrentDirectory)
	}

	files, err := ioutil.ReadDir(path)

	if err != nil {
		return errors.New(NoFilesInDirectory)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".exe") {
			execFile = file.Name()
		}
	}

	startThe := fmt.Sprintf(path + `\` + execFile)

	cmd := exec.Command(startThe)
	cmd.Run()

	return nil
}
