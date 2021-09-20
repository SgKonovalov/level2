package functions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
1)  Функция  ReadFromFile() – считывает текст из файла и в результате выдаёт срез строк (исходный срез.)
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
2) Функция  SortMain() – обрабатывает запрос пользователя
(а так же поддерживает ключи с, u + фильтрует по возрастанию и убыванию):
Принимая в качестве аргумента срез типа string (текст из файла), а так же сведения о наличии ключей
-n, -r, -c, -u, -b в команде пользователя (типа bool), производит обработку исходного текста, в зависимости от ключа:

* ключи -b и -c, поддерживаются в каждой комбинации:
	a) ключ -c: выводить в консоль информацию о сортировке текста;
	b) ключ -b: убирает пробелы из конца строки:
		- cортирует результирующий срез;
		- при помощи функции strings.TrimSpace убирает пробелы;
		- добавляет эти значения в промежуточный срез;
		- «обнуляет» итоговый срез;
		- добавляет значения из промежуточного среза в итоговый.

	2.1. Ключ -u:
		2.1.1. Создаём map, для хранения данных исходного среза;
		2.1.2. Перебираем исходный срез, и добавляем строки из исходного массива, в качестве ключа =>
		повторяющиеся строки исключаются;
		2.1.3. «Обнуляем» исходный срез;
		2.1.4. Перебираем map из п. 1.1.2 и добавляем значения ключа из неё в пустой исходный срез.

Сортировка по ключу -k и последующая обработка ключей -n, -r:

	2.2. Выборка колонки производится следующим образом:
		2.2.1. Наполняем срез sourceWithoutPoint строками из исходного файла (среза),
		убирая все знаки препинания из них (функция janitorOfPunctuation());
		2.2.2. Создаём регулярное выражение, применительно к запросу пользователя:
		в соответствии с номером указанной колонки (функция patternBuilder()) и записываем его
		в переменную resultPattern;
		2.2.3. Перебираем срез из п. 1.2.1 и добавляем в новый срез – forColumnData строки,
		совпавшие с регулярным выражением из п. 1.2.2;
		2.2.4. Проверяем длину полученного среза в п. 1.2.3. Если он == 0, а пользователь указал номер колонки > 1
		=> была указана либо последняя колонка, либо не существующая (слишком большое значение).
		В таком случае, в пустой срез из п. 1.2.3 добавляем все ПОСЛЕДНИЕ колонки из исходного среза (текста);
		2.2.5. Перебираем срезы: исходный (текст) и получившийся в пп. 1.2.3 или 1.2.4.
			2.2.5.1. Находим совпадения в элементах срезов и изменяем строки из исходного среза,
			добавляя к ним порядковый номер элемента из срезов п. 1.2.3/1.2.4.
			Добавление этих номеров позволит фильтровать элементы в ключах -n и -r;
			2.2.5.2. Заполняем получившимися значениями срез sliceWithNumber.
	1.3. Ключ -n:
		1.3.1. При помощи функции sort.Strings, сортируем строки из среза, полученного после отработки -k алгоритма;
		1.3.2. Добавляем в результирующий срез (на «выход») элементы из среза, полученного в п. 1.3.1,
		убирая первый символ (добавленные порядковый номера для сортировки).
	1.4. Ключ -n:
		1.4.1. при помощи функции sort.Slice, которой передаём в качестве аргумента исходный срез,
		полученный после отработки -k алгоритма и анонимную функцию, заменяющую элементы в этом срезе по формуле:
		предыдущий элемент за последующим, сортируем полученного после отработки -k алгоритма в обратном порядке;
		1.4.2. Добавляем в результирующий срез (на «выход») элементы из среза, полученного в п. 1.3.1,
		убирая первый символ (добавленные порядковый номера для сортировки).
	1.5. Ключ -h:
		1.5.1. Перебираем исходный срез, в ходе которого:
			1.5.1.1. При помощи regexp.MustCompile сравниваем строку с регулярным выражением,
			а именно, ищем строку, цифры в которой содержат суффикс;
			1.5.1.2. Если находим – добавляем их в промежуточный срез forSuffix;
		1.5.2. Запускаем цикл по количеству итераций == длине исходного среза,
		параллельно присваивая строке mainString значение элемента исходного среза.
			1.5.2.1. Внутри этого цикла перебираем срез forSuffix.
				- Если есть совпадения, то копируем исходный срез со следующими параметрами:
				в элементы ДО совпавшего элемента, копируем все элемент, после совпавшего;
				- Уменьшаем счётчик элементов на 1 и выходим из if.
		1.5.3. Сортируем срез forSuffix;
		1.5.4. Добавляем срез forSuffix в конец исходного среза.
*/
func SortMain(stringMassive []string, nkey, rkey, ckey, ukey, bkey, hkey bool, numOfColumn int) (result []string) {

	if ukey {

		var withoutSpace []string

		var uniq = make(map[string]bool)
		for _, uniqStr := range stringMassive {
			uniq[uniqStr] = true
		}

		stringMassive = nil

		for key, _ := range uniq {
			stringMassive = append(stringMassive, key)
		}

		if bkey {
			for _, withEndSpace := range stringMassive {
				withoutEndSpace := strings.TrimSpace(withEndSpace)
				withoutSpace = append(withoutSpace, withoutEndSpace)
			}
			result = nil
			result = append(result, withoutSpace...)
		}

		if ckey {
			fmt.Println("Text is sorted.")
		}

	}

	var forColumnData []string
	var sliceWithNumber []string

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

	intermSlice := lastSpacesTrimmer(forColumnData)

	sort.Strings(intermSlice)

	for _, sourceString := range stringMassive {
		for number, word := range intermSlice {
			if strings.Contains(sourceString, word) {
				sliceWithNumber = append(sliceWithNumber, fmt.Sprintf(strconv.Itoa(number)+sourceString))
			}
		}
	}

	if nkey {

		var withoutSpace []string

		sort.Strings(sliceWithNumber)

		for _, insert := range sliceWithNumber {
			result = append(result, insert[1:])
		}
		if bkey {
			for _, withEndSpace := range result {
				withoutEndSpace := strings.TrimSpace(withEndSpace)
				withoutSpace = append(withoutSpace, withoutEndSpace)
			}
			result = nil
			result = append(result, withoutSpace...)
		}

		if ckey {
			fmt.Println("Text is sorted.")
		}

		return result
	}

	if rkey {

		var withoutSpace []string

		sort.Slice(sliceWithNumber, func(prev, last int) bool {
			return sliceWithNumber[last] < sliceWithNumber[prev]
		})

		for _, insert := range sliceWithNumber {
			result = append(result, insert[1:])
		}

		if bkey {
			for _, withEndSpace := range result {
				withoutEndSpace := strings.TrimSpace(withEndSpace)
				withoutSpace = append(withoutSpace, withoutEndSpace)
			}
			result = nil
			result = append(result, withoutSpace...)
		}

		if ckey {
			fmt.Println("Text is sorted by reverse order.")
		}

		return result
	}

	if hkey {

		var intermSlise []string
		var forSuffix []string

		for _, toTheEnd := range stringMassive {

			numbers := regexp.MustCompile(`(?i)^\w+k\.\s`)

			if numbers.FindString(toTheEnd) != "" {
				forSuffix = append(forSuffix, toTheEnd)
			}
		}

		for nElem := 0; nElem < len(result); nElem++ {
			mainString := stringMassive[nElem]
			for _, rem := range forSuffix {
				if mainString == rem {
					stringMassive = append(result[:nElem], result[nElem+1:]...)
					nElem--
					break
				}
			}
		}

		sort.Strings(forSuffix)

		stringMassive = append(stringMassive, forSuffix...)

		if bkey {
			for _, withEndSpace := range result {
				withoutEndSpace := strings.TrimSpace(withEndSpace)
				intermSlise = append(intermSlise, withoutEndSpace)
			}
			result = nil
			result = append(stringMassive, intermSlise...)
		}

		if ckey {
			fmt.Println("Text is sorted by reverse order.")
		}

		return result
	}

	return result
}

/*
3) Функция  SortM() – фильтрует строки по месяцам, на основании наличии ключа
(а так же поддерживает ключи с, u + фильтрует по возрастанию и убыванию):
Принимая в качестве аргумента срез типа string (текст из файла), а так же сведения о наличии ключей
-n, -r, -c, -u в команде пользователя (типа bool), производит обработку исходного текста, в зависимости от ключа:

* ключи -b и -c, поддерживаются в каждой комбинации:
	a) ключ -c: выводить в консоль информацию о сортировке текста;
	b) ключ -b: убирает пробелы из конца строки:
		- cортирует результирующий срез;
		- при помощи функции strings.TrimSpace убирает пробелы;
		- добавляет эти значения в промежуточный срез;
		- «обнуляет» итоговый срез;
		- добавляет значения из промежуточного среза в итоговый.

	3.1. Ключ -u:
		2.1.1. Создаём map, для хранения данных исходного среза;
		2.1.2. Перебираем исходный срез, и добавляем строки из исходного массива, в качестве ключа =>
		повторяющиеся строки исключаются;
		2.1.3. «Обнуляем» исходный срез;
		2.1.4. Перебираем map из п.  1.1.2 и добавляем значения ключа из неё в пустой исходный срез.
	3.2. Ключ -n: перебираем исходный срез и если его слово его строки совпадает с названием месяца,
	добавляем эту строку в срез, в порядке возрастания месяцев (январь – декабрь);
	3.3. Ключ -r: перебираем исходный срез и если его слово его строки совпадает с названием месяца,
	добавляем эту строку в срез, в порядке убывания месяцев (декабрь – январь)
*/

func SortByM(stringMassive []string, nkey, rkey, ckey, bkey bool) (result []string) {

	if nkey {

		var withoutSpace []string
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "January") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "February") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "March") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "April") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "May") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "June") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "July") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "August") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "September") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "October") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "November") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "December") {
				result = append(result, thirdColumn)
			}
		}

		if bkey {
			for _, withEndSpace := range stringMassive {
				withoutEndSpace := strings.TrimSpace(withEndSpace)
				withoutSpace = append(withoutSpace, withoutEndSpace)
			}
			result = nil
			result = append(result, withoutSpace...)
		}

		if ckey {
			fmt.Println("Text is sorted.")
		}
		return result
	}

	if rkey {
		var withoutSpace []string
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "December") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "November") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "October") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "September") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "August") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "July") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "June") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "May") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "April") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "March") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "February") {
				result = append(result, thirdColumn)
			}
		}
		for _, thirdColumn := range stringMassive {
			if strings.Contains(thirdColumn, "January") {
				result = append(result, thirdColumn)
			}
		}

		if bkey {
			for _, withEndSpace := range stringMassive {
				withoutEndSpace := strings.TrimSpace(withEndSpace)
				withoutSpace = append(withoutSpace, withoutEndSpace)
			}
			result = nil
			result = append(result, withoutSpace...)
		}

		if ckey {
			fmt.Println("Text is sorted.")
		}
		return result
	}

	return result
}
