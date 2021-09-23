package command

import (
	"regexp"
	"strconv"
	"strings"
)

/*
1) Функция LookForTimeout() – получая в качестве аргумента команду, введённую пользователем,
расшифровывает её и в результате выдаёт информацию о наличии опции timeout (значение типа bool)
и о длительности этого timeout (значение типа int).
	1.1. Перебирает строку команды и убираем из неё все пробелы;
	1.2. Добавляет строки без пробелов в промежуточный срез(array);
	1.3. Перебираем срез array. Если находим слово timeout:
		1.3.1. Присваиваем bool-значению, указывающему на наличие опции timeout – true;
		1.3.2. При помощи регулярного выражения, ищем длину timeout,
		указанного пользователем и переводим это число в тип string;
		1.3.3. При помощи strconv.Atoi, переводим длину timeout в int.
	1.4. Возвращаем полученный в п. 1.3.1 bool и в п. 1.3.3 int.
*/

func LookForTimeout(command string) (hasTimeout bool, timeout int) {

	var intoArray string
	var array []string
	var timeoutWithTime string

	for _, withoutSpace := range command {
		if withoutSpace != ' ' {
			intoArray += string(withoutSpace)
		}
		if withoutSpace == ' ' {
			if intoArray != "" {
				array = append(array, intoArray)
				intoArray = ""
			}
		}
	}
	if intoArray != "" {
		array = append(array, intoArray)
	}

	for _, lookForTO := range array {
		if strings.Contains(lookForTO, "timeout") {
			hasTimeout = true
			timeoutWithTime = lookForTO

		}
	}

	if hasTimeout {
		howLong := regexp.MustCompile("[0-9]+")
		interm := howLong.FindAllString(timeoutWithTime, -1)
		for _, giveTime := range interm {
			timeout, _ = strconv.Atoi(giveTime)
		}
	}

	return hasTimeout, timeout
}
