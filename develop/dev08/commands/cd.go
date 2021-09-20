package commands

import (
	"errors"
	"fmt"
	"os"
)

/*
Функция CD, осуществляет смену текущей директории на директорию, указанную пользователем.
Получает в качестве аргумента путь к директории, в которую нужно перейти.
С помощью os.Chdir подключаемся к shell и передаём в качестве аргумента путь к файлу из аргумента.
Если указанный путь к папке не существует – выдаём ошибку: cantFoundCurrentDirectory,
в противном случае, выдаём сообщение об успешной смене директории.
*/

func CD(pathToDirectory string) (path string, err error) {

	err = os.Chdir(pathToDirectory)
	if err != nil {
		return "", errors.New(cantFoundCurrentDirectory)
	}

	path = fmt.Sprintf("Directory changed at %s successful", pathToDirectory)

	return path, nil
}
