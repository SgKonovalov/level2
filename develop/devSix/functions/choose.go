package functions

import (
	"errors"

	"exeTwo.devThree/control"
)

/*
Функция ChooseOption() – координирует работу с командами и вызывает нужную функцию из файла funcSet.go,
в соответствии с введёнными ключами.
*/

func ChooseOption(fkey, dkey, skey bool, numOfColumn int) (result []string, err error) {

	if !fkey && !dkey && !skey {
		return nil, errors.New(noCommands)
	}

	if fkey && numOfColumn <= 0 {
		return nil, errors.New(wrongColNum)
	}

	if fkey && skey {
		return nil, errors.New(wrongKeysSetSF)
	}

	if fkey && numOfColumn > 0 {
		if dkey {
			return nil, errors.New(wrongKeysSetDF)
		} else {
			result = SortByfKey(ReadFromFile(control.ReadFileNameFromConsole()), numOfColumn)
			return result, nil
		}
	}

	if skey {
		result = SortBysKey(ReadFromFile(control.ReadFileNameFromConsole()))
		if result == nil {
			return nil, errors.New(emptyResult)
		}
		return result, nil
	}

	if dkey && !fkey {
		result = SortBydKey(ReadFromFile(control.ReadFileNameFromConsole()), control.ReadNewSeparator())
		if result == nil {
			return nil, errors.New(emptyResult)
		}
		return result, nil
	}

	return result, nil
}
