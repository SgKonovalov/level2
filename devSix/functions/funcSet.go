package functions

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
1) Функция ReadCommandFromConsole() - читает введённую пользователем команду и возвращает её текст;
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
2) Функция SortByfKey() - получает срез типа string и номер колонки, которую нужно найти.
Выборка колонки производится следующим образом:
	2.1. Наполняем срез sourceWithoutPoint строками из исходного файла (среза),
	убирая все знаки препинания из них (функция janitorOfPunctuation());
	2.2. Создаём регулярное выражение, применительно к запросу пользователя:
	в соответствии с номером указанной колонки (функция patternBuilder()) и записываем его в переменную resultPattern;
	2.3. Перебираем срез из п. 2.1 и добавляем в новый срез – forColumnData строки, совпавшие
	с регулярным выражением из п. 2.2;
	2.4. Проверяем длину полученного среза в п. 1.2.3. Если он == 0, а пользователь указал номер колонки > 1
	=> была указана либо последняя колонка, либо не существующая (слишком большое значение).
	В таком случае, в пустой срез из п. 2.3 добавляем все ПОСЛЕДНИЕ колонки из исходного среза (текста);
	2.5. В итоговый срез добавляем все значения, убирая концевые и начальные пробелы, при помощи функции
	lastSpacesTrimmer().
*/

func SortByfKey(stringMassive []string, numOfColumn int) (result []string) {

	if numOfColumn <= 0 {
		return nil
	}

	var forColumnData []string

	sourceWithoutPoint := janitorOfPunctuation(stringMassive)

	resultPattern := patternBuilder(numOfColumn)

	for _, word := range sourceWithoutPoint {

		forColumnData = append(forColumnData, resultPattern.FindStringSubmatch(word)...)
	}

	if numOfColumn > 1 && len(forColumnData) == 0 {
		resultPattern = regexp.MustCompile(`\w+$`)
		for _, word := range stringMassive {
			forColumnData = append(forColumnData, resultPattern.FindStringSubmatch(word)...)
		}
	}

	result = lastSpacesTrimmer(forColumnData)

	return result
}

/*
3) Функция SortBydKey() - получает срез типа string и символ типа string, которым нужно заменить разделитель.
	3.1. Проверяет длину строки полученного разделителя, если она > 1 => возвращаем пустой срез.
	3.2. Переводим строку разделителя в срез типа runeб перебираем этот срез и присваиваем значение rune
	разделителя переменной changeRune.
	3.3. Перебираем исходный срез (полученный в качестве аргумента):
		3.3.1. Переводим его строки в срез типа rune;
		3.3.2. Перебираем каждый из этих срезов;
		3.3.3. Ели в процессе переборки, встречаем rune типа TAB, заменяем её на changeRune;
		3.3.4. Добавляем строку из п 3.3.3 в промежуточный срез (intrmArr).
	3.4. Добавляем в итоговый срез, срез из п. 3.3.4.
*/

func SortBydKey(stringMassive []string, sign string) (result []string) {

	if len(sign) > 1 {
		return nil
	}

	var changeRune rune
	var intrmArr []rune

	signToRune := []rune(sign)
	for _, signIsRune := range signToRune {
		changeRune = signIsRune
	}

	for _, word := range stringMassive {
		allTextToRune := []rune(word)
		for _, findSep := range allTextToRune {
			if findSep == 32 {
				findSep = changeRune
			}
			intrmArr = append(intrmArr, findSep)
		}
	}

	result = append(result, string(intrmArr))
	return result
}

/*
4) Функция SortBysKey() - получает срез типа string.
	4.1. Перебираем исходный срез (полученный в качестве аргумента);
	4.2. Если в строке попадается совпадения с rune типа TAB – добавляем эту строку в итоговый срез;
	4.3. Добавляем в итоговый срез, срез из п. 4.2.
*/

func SortBysKey(stringMassive []string) (result []string) {

	for _, word := range stringMassive {
		if strings.ContainsRune(word, 32) {
			result = append(result, word)
		}
	}
	return result
}
