/*
ПРИЛОЖЕНИЕ:
_________________________________________________________________________________________________________
1) Определяем переменные:
exceptionStringFromNumbers – текст ошибки для вывода в консоль. Ситуация: строка состоит только из чисел;
exceptionEmptyString – текст ошибки для вывода в консоль. Ситуация: задана пустая строка.
Цель определения этих переменных: предупреждение «error strings should not be capitalized»,
т.е. нельзя писать текст с заглавной буквы.

2) Функция translateString().
	2.1. Принимает в качестве аргумента исходную строку, которую нужно распаковать;
	2.2. В результате работы функции получаем либо нужную строку и ошибку.
	2.3. Определяем перменные:
		2.3.1. tsRune типа rune. Цель – определить все символы в строке;
		2.3.2. getTheString типа strings.Builder – чтобы далее строить строки с его помощью.
	2.4. В первом блоке if – else if проводим проверку на «пустую» строку или строку,
	состоящую только из цифр. Если возникает одна из вышеперечисленных проблем – выходим из функции,
	возвращаем «пустую» строку и ошибку. В противном случае, «идём» дальше.
	2.5. С помощью цикла for, итерируемся по символам строки:
		2.5.1. if unicode.IsDigit(letter) – проверяем является ли символ цифрой;
		2.5.2. Если «ЯВЛЯЕТСЯ»:
			2.5.2.1. Переводим эти символы в int;
			2.5.2.2. С помощью strings.Repeat и tsRune создаём новую строку, дублируя символ столько раз,
			сколько указано в count.
		2.5.3. Если «НЕ ЯВЛЯЕТСЯ»:
			2.5.3.1. Проверяем на наличие в строке символов \\.
			Если таких нет – создаём строку с помощью getTheString и метода WriteRune.
			2.5.3.2. Если строка без цифр и вышеуказанных символов – выводим её не изменяя.
	2.6. В блоке return с помощью getTheString – выдаём получившуюся строку и ошибку.

ТЕСТЫ:
_________________________________________________________________________________________________________
1) При помощи функции TestTranslateString() – тестируем основной функционал.
	1.1. В map задаём строку – ключ и верный результат – значение;
	1.2. Запускаем тестируемую функцию translateString() по количеству значений в map из п. 1.1;
	1.3. Сверяем результаты: если результат, полученный после работы тестируемой функции совпал с верным результатом
	из map (п. 1.1) – тест пройден. Если нет – выдаём ошибку.
2) При помощи функции TestWrongStringsTranslateString() – тестируем функционал по отбору «неверных» строк (пустые или только из чисел).
	2.1. Создаём строку для тестирования;
	2.2. Запускаем функцию translateString().
	2.3. Если в результате работы, функция translateString() «выдала» пустую строку и ошибку – тест пройден.
*/

package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var (
	exceptionStringFromNumbers = "Input string contains numbers only"
	exceptionEmptyString       = "Input empty string"
)

func translateString(sourceString string) (resultString string, err error) {

	var (
		tsRune       rune
		getTheString strings.Builder
	)

	if _, err := strconv.Atoi(string(sourceString)); err == nil {
		return "", errors.New(exceptionStringFromNumbers)
	} else if sourceString == "" {
		return "", errors.New(exceptionEmptyString)
	}

	for _, letter := range sourceString {

		if unicode.IsDigit(letter) {

			if count, err := strconv.Atoi(string(letter)); err == nil {
				intermResult := strings.Repeat(string(tsRune), count-1)
				getTheString.WriteString(intermResult)
			}

		} else {
			if !(string(letter) == "\\" && string(letter) != "\\") {

				getTheString.WriteRune(letter)
			}
			tsRune = letter
		}

	}

	return getTheString.String(), err
}

func main() {
	parseWord, err := translateString("qwe\\5")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(parseWord)
	}
}
