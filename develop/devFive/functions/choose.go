package functions

import (
	"errors"

	"exeTwo.devThree/control"
)

/*
Текст ошибок:
wrongStrNum - не указано количество строк в ключах А, В или С.
emptyResult - в результате работы возвращён "пустой" срез. Проблема в запросе.
allKeys - если выбрано больше 1 ключа.
*/

var wrongStrNum = "Selected wrong numbers of string in 'A/B/C' keys. Please insert column number in command, like 'A2'"
var emptyResult = "Received the empty result. Please, check your command!"
var allKeys = "Your can't choose more than one key! Please using ONE ONLY."

/*
Функция ChooseOption() – координирует работу с командами и вызывает нужную функцию из файла funcSet.go,
в соответствии с введёнными ключами.
*/

func ChooseOption(Akey, Bkey, Ckey, ckey, ikey, vkey, Fkey, nkey bool, numOfStrings int) (result interface{}, err error) {

	var checkKeys []bool
	if Akey {
		checkKeys = append(checkKeys, Akey)
	}
	if Bkey {
		checkKeys = append(checkKeys, Bkey)
	}
	if Ckey {
		checkKeys = append(checkKeys, Ckey)
	}
	if ckey {
		checkKeys = append(checkKeys, ckey)
	}
	if ikey {
		checkKeys = append(checkKeys, ikey)
	}
	if vkey {
		checkKeys = append(checkKeys, vkey)
	}
	if Fkey {
		checkKeys = append(checkKeys, Fkey)
	}
	if nkey {
		checkKeys = append(checkKeys, nkey)
	}

	if len(checkKeys) > 1 {
		return nil, errors.New(allKeys)
	}

	if ckey {
		result = SortBycKey(ReadFromFile(control.ReadFileNameFromConsole()))
		return result, nil
	}

	if Fkey {
		result = SortByFKey(ReadFromFile(control.ReadFileNameFromConsole()), control.ReadTextFromConsole())
		return result, nil
	}

	if ikey {
		result = SortByiKey(ReadFromFile(control.ReadFileNameFromConsole()), control.ReadTextFromConsole())
		return result, nil
	}

	if vkey {
		result = SortByvKey(ReadFromFile(control.ReadFileNameFromConsole()), control.ReadTextFromConsole())
		return result, nil
	}

	if nkey {
		result = SortBynKey(ReadFromFile(control.ReadFileNameFromConsole()), control.ReadTextFromConsole())
		return result, nil
	}

	if Akey {

		if numOfStrings <= 0 {
			return nil, errors.New(wrongStrNum)
		} else {
			result = SortByAKey(ReadFromFile(control.ReadFileNameFromConsole()), control.ReadTextFromConsole(), numOfStrings)
			return result, nil
		}
	}

	if Bkey {

		if numOfStrings <= 0 {
			return nil, errors.New(wrongStrNum)
		} else {
			result = SortByBKey(ReadFromFile(control.ReadFileNameFromConsole()), control.ReadTextFromConsole(), numOfStrings)
			if err != nil {
				return nil, errors.New(emptyResult)
			}
		}
	}

	if Ckey {

		if numOfStrings <= 0 {
			return nil, errors.New(wrongStrNum)
		} else {
			result = SortByCKey(ReadFromFile(control.ReadFileNameFromConsole()), control.ReadTextFromConsole(), numOfStrings)
			return result, nil
		}
	}

	if result == nil {
		return nil, errors.New(emptyResult)
	}

	return result, nil
}
