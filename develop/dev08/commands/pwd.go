package commands

import (
	"errors"
	"os"
)

/*
Функция PWD() – при помощи os.Getwd получает путь к текущей директории
и возвращает его в return в качестве значения типа string.
*/

func PWD() (path string, err error) {

	path, err = os.Getwd()
	if err != nil {
		return "", errors.New(cantFoundCurrentDirectory)
	}

	return path, nil
}
