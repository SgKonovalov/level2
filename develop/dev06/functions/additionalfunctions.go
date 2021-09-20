package functions

import (
	"bytes"
	"regexp"
	"strings"
)

/*
1) Функция janitorOfPunctuation() – убирает все знаки препинания из исходной строки.
Цель: соответствие универсальному регулярному выражению `\w+\s`.
Алгоритм работы функции:
	1.1. На вход получает исходный текст;
	1.2. Перебирает исходный текст:
		1.2.1. Если находит в строке исходного текста соответствие регулярному выражению – знак препинания
		(для примера `\.`), при помощи функции ReplaceAllString() – заменяет её на пустой.
	1.3. Добавляет полученные строки в результирующий срез, на «выход».
*/

func janitorOfPunctuation(stringMassive []string) (sourceWithoutPoint []string) {

	for _, word := range stringMassive {
		withoutPoint := regexp.MustCompile(`\.`)
		noPoint := withoutPoint.ReplaceAllString(word, "")
		sourceWithoutPoint = append(sourceWithoutPoint, noPoint)
	}
	return sourceWithoutPoint
}

/*
2) Функция patternBuilder() – строит регулярное выражение для дальнейшего поиска строк в соответствии с ним.
Алгоритм работы функции:
	2.1. На вход получает номер колонки, в рамках которой пользователь хочет произвести операцию;
	2.2. Запускает цикл по количеству итераций == номеру колонки.
	В рамках этого цикла, при помощи bytes.Buffer{} «строим» регулярное выражение,
	добавляя к нему регулярное выражение `\w+\s` по количеству итераций «минус» 1.
	Пример: пользователь хочет произвести сортировку по 4 колонке => в рамках этого цикла,
	получаем регулярное выражение вида: `\w+\s\w+\s\w+\s`.
	2.3. Строим итоговый паттерн, добавляя к нему первую колонку.
	ВАЖНО: т.к. итерация в п. 1.2 начинается с 1 => если пользователь задаст первую колонку, то результирующее регулярное выражение будет == `^\w+\s`.
*/

func patternBuilder(numOfColumn int) (resultPattern *regexp.Regexp) {

	var forPattern []string

	firstColumn := `^\w+\s`
	afterTo := `\w+\s`

	for i := 1; i < numOfColumn; i++ {
		forPattern = append(forPattern, afterTo)
	}

	buffer := bytes.Buffer{}

	for _, patternForString := range forPattern {
		buffer.WriteString(patternForString)
	}

	resultPattern = regexp.MustCompile(firstColumn + buffer.String())

	return resultPattern
}

/*
3) Функция lastSpacesTrimmer() – убирает крайний пробел из строк, содержащий только слово из колонки строки.
Цель: в промежуточный срез попадут все слова и колонок, ДО которой пользователь осуществляет выборку =>
нам понадобится последний его элемент, а самым простым способом поиска этого слова является
использование регулярного выражения `\w+$`.
Алгоритм работы функции:
3.1. На вход получает срез со словами из колонок;
3.2. При помощи функции TrimSpace() – обрезает пробелы;
3.3. Добавляет получившиеся значения в итоговый срез.
*/

func lastSpacesTrimmer(forColumnData []string) (intermSlice []string) {

	for i := 0; i < len(forColumnData); i++ {
		withoutLastSpace := strings.TrimSpace(forColumnData[i])
		sec := regexp.MustCompile(`\w+$`)
		intermSlice = append(intermSlice, string(sec.Find([]byte(withoutLastSpace))))
	}
	return intermSlice
}
