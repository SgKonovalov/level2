package functions

import (
	"errors"

	"exeTwo.devThree/control"
)

/*
1) Переменная var wrongKKey = "The column is not selected. Please insert column number in command, like 'sort -k1'"
– это текст выдаваемой ошибки, если пользователь не выбрал строку номер колонки для обработки.
2) Переменная var emptyResult = "Received the empty result. Please, check your command!"
– это текст выдаваемой ошибки, если в ответе на запрос пользователя возвращён «пустой» результат.
3) Переменная var wrongMonthFilter = "You can't choose key 'M' without key 'n' or 'r'. Please change your command, like 'sort -k1Mn'"
– это текст выдаваемой ошибки, если в запросе, пользователь не указал порядок сортировки по месяцам.
*/

var wrongKKey = "The column is not selected. Please insert column number in command, like 'sort -k1'"
var emptyResult = "Received the empty result. Please, check your command!"
var wrongMonthFilter = "You can't choose key 'M' without key 'n' or 'r'. Please change your command, like 'sort -k1Mn'"

/*
3) Функция ChooseOption() – координирует работу с командами и вызывает нужную функцию из файла funcSet.go,
в соответствии с введёнными ключами. Так, в качестве аргументов ей передаётся «комплект» значений типа bool,
соответствующий каждому ключу. Если значений нет для ключа k или отсутствует номер колонки в этом ключе =>
выходим из функции с ошибкой wrongKKey. В зависимости от номера колонки, вызываем обработчик,
передавая ему аргументы для запроса соответствующих функций.
*/

func ChooseOption(kkey, nkey, rkey, ukey, ckey, Mkey, bkey, hkey bool, numOfColumn int) (result []string, err error) {

	if !kkey || numOfColumn <= 0 {
		return nil, errors.New(wrongKKey)
	}

	if !Mkey {
		text := control.ReadFileNameFromConsole()
		result = SortMain(ReadFromFile(text), nkey, rkey, ckey, ukey, bkey, hkey, numOfColumn)
		if len(result) == 0 {
			return nil, errors.New(emptyResult)
		}
		return result, nil
	}

	if Mkey && !nkey && !rkey {
		return nil, errors.New(wrongMonthFilter)
	}

	if Mkey {
		text := control.ReadFileNameFromConsole()
		result = SortByM(ReadFromFile(text), nkey, rkey, ckey, bkey)
		if len(result) == 0 {
			return nil, errors.New(emptyResult)
		}
		return result, nil
	}

	return result, nil
}
