package functions

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

/*
1) Функция ReadFromFile() – получая в качестве аргумента имя файла, читает текст из него текст,
переводя его в срез типа string и возвращает полученный результат.
*/

func ReadFromFile(fileName string) (stringMassive []string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringMassive = append(stringMassive, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return stringMassive
}

/*
2) Функция SortBycKey() – возвращает значения типа int.
Это есть длина среза, полученного от функции ReadFromFile().
*/

func SortBycKey(stringMassive []string) (stringsNumber int) {
	return len(stringMassive)
}

/*
3) Функция SortByFKey() – получает срез типа string и часть текста, который нужно найти.
В случае совпадения выдаёт в результате фрагмент из текста. Если такого фрагмента нет – выдаёт ошибку:
"Desired fragment or text is undetected".
*/

func SortByFKey(stringMassive []string, searchingText string) (result string) {
	for _, finded := range stringMassive {

		ok, _ := regexp.MatchString(searchingText, finded)
		if ok {
			result = finded
		}
	}
	if result == "" {
		result = "Desired fragment or text is undetected"
	}
	return result
}

/*
4) Функция SortByiKey() – выполняет тот же самый алгоритм, что и SortByFKey(),
приводя текст из запроса пользователя и текст из файла к строчный буквам.
В случае совпадения выдаёт в результате фрагмент из текста.
Если такого фрагмента нет – выдаёт ошибку: "Desired fragment or text is undetected".
*/

func SortByiKey(stringMassive []string, searchingText string) (result string) {

	for _, finded := range stringMassive {

		ok := strings.Contains(strings.ToLower(searchingText), strings.ToLower(finded))
		if ok {
			result = finded
		}
	}
	if result == "" {
		result = "Desired fragment or text is undetected"
	}
	return result
}

/*
5) Функция SortByvKey() – принимает в качестве аргументов срез типа string, содержащий текст из файла,
а так же string – фрагмент текста, введённый пользователем,
для его исключения из результатов (вместо совпадения, исключать).
Функция:
	5.1.. Перебирает полученный срез, и добавляет из него строки в результирующий массив.
	5.2. Если строка из исходного среза == строке из текста, такая строка исключается
	и не добавляется в результирующий срез.
	5.3. Если все строки == тексту, введённому пользователем,
	то на выходе выдаём ошибку "Other strings are undetected".
	Чтобы вывести такую ошибку, сравниваем длину результирующего срезе с 0. Если они ==,
	то выдаём вышеуказанную ошибку.
	5.4. Возвращаем результирующий срез.
*/

func SortByvKey(stringMassive []string, searchingText string) (result []string) {

	for _, finded := range stringMassive {

		ok := strings.Contains(searchingText, finded)
		if !ok {
			result = append(result, finded)
		}
	}
	if len(result) == 0 {
		result = append(result, "Other strings are undetected")
	}
	return result
}

/*
6) Функция SortBynKey() – принимает в качестве аргументов срез типа string,
содержащий текст из файла, а так же string – фрагмент текста, введённый пользователем.
Функция:
Перебирает полученный срез. Если строка из исходного среза == фрагменту текста,
введённого пользователем, возвращаем номер строки как элемента исходного среза +1.
*/

func SortBynKey(stringMassive []string, searchingText string) (result int) {

	for stringNumber, finded := range stringMassive {

		ok := strings.Contains(searchingText, finded)
		if ok {
			result = stringNumber + 1
		}
	}
	return result
}

/*
7) Функция SortByАKey() – принимает в качестве аргументов срез типа string,
содержащий текст из файла, а так же string – фрагмент текста,
введённый пользователем и количество строк, который хочет отобрать пользователь.
	7.1. Перебирает полученный срез.
	7.2. Если строка из исходного среза == фрагменту текста, введённого пользователем,
	запускаем новый цикл, от 1 до «количества строк, который хочет отобрать пользователь» + 1,
	в котором в результирующий срез добавляем строки из исходного массива,
	которые следуют за указанной пользователем.
	7.3. Возвращаем результирующий срез.
*/

func SortByAKey(stringMassive []string, searchingText string, numOfStrings int) (result []string) {

	for stringNumber, finded := range stringMassive {

		if strings.Contains(searchingText, finded) {
			for i := 1; i < numOfStrings+1; i++ {
				result = append(result, stringMassive[stringNumber+i])
			}
		}

	}
	return result
}

/*
8) Функция SortByВKey() – принимает в качестве аргументов срез типа string,
содержащий текст из файла, а так же string – фрагмент текста, введённый пользователем и количество строк,
который хочет отобрать пользователь.
	8.1. Перебирает полученный срез.
	8.2. Если строка из исходного среза == фрагменту текста, введённого пользователем,
	копируем в результирующий срез элементы из исходного массива, по формуле:
	«НАЧАЛО» с номера совпавшего элемента «минус» количество строк указанных пользователем
	«ДО» номера совпавшего элемента(не включая его).
	8.3. Сортируем результирующий срез, меняя местами элементы.
	8.4. Возвращаем результирующий срез.
*/

func SortByBKey(stringMassive []string, searchingText string, numOfStrings int) (result []string) {

	for stringNumber, finded := range stringMassive {

		if strings.Contains(searchingText, finded) {

			result = stringMassive[stringNumber-numOfStrings : stringNumber]

		}
	}
	sort.Slice(result, func(prev, last int) bool {
		return result[last] < result[prev]

	})
	return result
}

/*
9) Функция SortByСKey() – принимает в качестве аргументов срез типа string,
содержащий текст из файла, а так же string – фрагмент текста, введённый пользователем и количество строк,
который хочет отобрать пользователь.
	9.1. Добавляем в результирующий срез, срез полученный после вызова функции SortByВKey();
	9.2. Добавляем в результирующий срез, срез полученный после вызова функции SortByАKey().
*/

func SortByCKey(stringMassive []string, searchingText string, numOfStrings int) (result []string) {

	result = append(result, SortByBKey(stringMassive, searchingText, numOfStrings)...)
	result = append(result, SortByAKey(stringMassive, searchingText, numOfStrings)...)

	return result
}
