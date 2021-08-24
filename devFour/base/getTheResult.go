package base

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

var notMaches = "Entering array has no anagrams of vocabularies."
var smallVocabulariesList = "Returned only one or zero vocabularies."

/*
Функция GetTheResult() – принимает в качестве аргумента 2 значения типа bool,
которые отражаются совпадения с анаграммами в словаре.
В результате работы выдаёт готовую map с ключом типа string (любое слово из анаграммы)
и значением типа среза string == набору всех анаграмм.
По работе функции:
1) Если оба переданных значения – false, выходим из функции,
выдавая ошибку «Entering array has no anagrams of vocabularies».
2) При наличии совпадения с одним или несколькими анаграммами – сортируем готовый срез анаграмм,
в качестве ключа устанавливаем random значение и записываем в итоговую map.
3) Проверяем длину итоговой map, если она < 2 => в результате выдаём ошибку:
«Returned only one or zero vocabularies»; в противном случае – выдаём итоговую map.
*/

func GetTheResult(includeFirst, includeSecond bool) (map[string][]string, error) {

	if !(includeFirst) && !(includeSecond) {
		return nil, errors.New(notMaches)
	}

	var result = make(map[string][]string)

	if includeFirst {
		sort.Strings(FirstLib)
		key := FirstLib[(rand.Intn(len(FirstLib)))]
		result[key] = FirstLib
	}
	if includeSecond {
		sort.Strings(SecondLib)
		key := SecondLib[(rand.Intn(len(SecondLib)))]
		result[key] = SecondLib
	}

	if len(result) < 2 {
		return nil, errors.New(smallVocabulariesList)
	}

	return result, nil
}

/*
Функция PrintResult() – получает итоговую map и просто выводит её данные в консоль.
*/

func PrintResult(forPrint map[string][]string) {
	for key, value := range forPrint {
		fmt.Printf("Key is: %s, vocabularies: %v\n", key, value)
	}
}
