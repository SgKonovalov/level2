package base

import (
	"reflect"
	"sort"
	"strings"
)

/*
Работа с паттернами:

1) Инициализируем паттерны (любое слово из словаря). В дальнейшем будем их использовать для сравнения:
var PatternByFirstSet = []rune("пятак")
var PatternBySecondSet = []rune("листок")

2) Готовые срезы для словаря:
var FirstLib = []string{"пятак", "пятка", "тяпка"}
var SecondLib = []string{"листок", "слиток", "столик"}

3) Устанавливаем значения результата сравнения полученых даных и анаграмм в false:
var ForFirst = false
var ForSecond = false
*/

var PatternByFirstSet = []rune("пятак")
var PatternBySecondSet = []rune("листок")

var FirstLib = []string{"пятак", "пятка", "тяпка"}
var SecondLib = []string{"листок", "слиток", "столик"}

var ForFirst = false
var ForSecond = false

/*
4) Функция PreparedingPatterns() – принимает в качестве аргументы все наборы анаграмм из словаря, возвращая 2 СОРТИРОВАННЫХ среза типа int, которые составляют 2 паттерна PatternByFirstSet и PatternBySecondSet:
	4.1. Перебираем полученные аргументы, переволдя каждую rune в int и добавляем полученное значение в срез «для выдачи».
	4.2. Сортируем этот срез.
*/

func PreparedingPatterns(patternByFirstSet, patternBySecondSet []rune) (firstPattern, secondPattern []int) {

	for _, fptByRune := range patternByFirstSet {
		fptByInt := int(fptByRune - '0')
		firstPattern = append(firstPattern, fptByInt)
	}
	sort.Ints(firstPattern)

	for _, sptByRune := range patternBySecondSet {
		sptByInt := int(sptByRune - '0')
		secondPattern = append(secondPattern, sptByInt)
	}
	sort.Ints(secondPattern)

	return firstPattern, secondPattern
}

/*
5) Функция Matcher() – принимает в качестве аргумента срез типа string (данные полученные для проверки), в результате возвращает 2 значения типа bool (соответствие паттернам):
	5.1. Перебираем полученный срез:
		5.1.1. Переводим все буквы слова из полученного срез в строчные;
		5.1.2. Переводим слова в []rune;
		5.1.3. Переводим полученные rune в int, добавляем всё это в срез и сортируем;
	5.2. Вызываем функцию PreparedingPatterns(), получая срезы с int, сформированные из паттернов. Сравниваем эти срезы, если нашли совпадения – присваиваем соответствующему паттерну bool значение true.
*/

func Matcher(source []string) (ForFirst, ForSecond bool) {

	var result []rune
	var sd int

	for _, word := range source {
		word = strings.ToLower(word)
		result = []rune(word)
		var getTheInts []int
		for _, das := range result {
			sd = int(das - '0')
			getTheInts = append(getTheInts, sd)
		}
		sort.Ints(getTheInts)
		firstPattern, secondPattern := PreparedingPatterns(PatternByFirstSet, PatternBySecondSet)
		if reflect.DeepEqual(firstPattern, getTheInts) {
			ForFirst = true
		}
		if reflect.DeepEqual(secondPattern, getTheInts) {
			ForSecond = true
		}
	}
	return ForFirst, ForSecond
}
