package pkg

import (
	"encoding/json"
	"time"
)

/*
1) Функция ShowEventByDay() – принимая в качестве аргумента указатель на объект типа CalendarModel и его Хранилища,
делает выборку событий в указанный день:
	1.1. Перебираем Хранилище. Если находим событие, заданное пользователем,
	добавляем его значение в итоговый срез.
	1.2. Переводим результат в JSON и выдаём оба: срез и JSON.
*/

func ShowEventByDay(date CalendarModel, CalendarStorage []*CalendarModel) (result []CalendarModel, inJSON []byte) {

	for _, event := range CalendarStorage {
		if (event.Day == date.Day) && (event.Month == date.Month) && (event.Year == date.Year) {
			result = append(result, *event)
		}
	}

	inJSON, _ = json.Marshal(result)

	return result, inJSON
}

/*
2) Функция ShowEventByMonth() – принимая в качестве аргумента указатель на объект типа CalendarModel и его Хранилища,
делает выборку событий в указанный месяц:
	2.1. Перебираем Хранилище. Если находим события, по месяцу и году совпадающие
	со введёнными пользователем параметрами – сохраняем эти события в итоговый срез.
	2.2. Переводим результат в JSON и выдаём оба: срез и JSON.
*/

func ShowEventByMonth(date CalendarModel, CalendarStorage []*CalendarModel) (result []CalendarModel, inJSON []byte) {

	for _, event := range CalendarStorage {
		if (event.Month == date.Month) && (event.Year == date.Year) {
			result = append(result, *event)
		}
	}

	inJSON, _ = json.Marshal(result)

	return result, inJSON
}

/*
3) Функция ShowEventByWeek() – принимая в качестве аргумента указатель на объект типа CalendarModel и его Хранилища,
делает выборку событий за неделю (указанный пользователем «минус» 7 дней):
	3.1. Перебираем Хранилище. Если находим события, по месяцу и году совпадающие
	со введёнными пользователем параметрами – сохраняем дату отсчёта в переменную weekCount, по формуле день,
	указанный пользователем «минус» 6 (ПРИМЕР: пользователь вводит дату: 07.08.2021,
	в weekCount сохраняем день: 7 – 6 = 1).
	3.2. Далее работа функции зависит от результат weekCount:
		3.2.1. Если weekCount > 7 или == 0 (т.е. 7 и последующие числа месяца):
			3.2.1.1. Вызываем функцию ShowEventByMonth() и сохраняем полученные результаты в промежуточный срез
			otherDays;
			3.2.1.2. Перебираем срез otherDays:
			-если день > weekCount и день <= дню, заданному пользователем, добавляем эти результаты в итоговой срез.
		3.2.2. Если weekCount < 7 или != 0 (т.е. до 7 числа месяца):
			3.2.2.1. Инициализируем объект-заглушку типа CalendarModel
			3.2.2.2. Находим дни, подпадающие под недельный отрезок из ПРЕДЫДУЩЕГО месяца:
			- сохраняем в переменную otherDaysInPrevMonth срез,
			полученный в результате работы функции ShowEventByMonth();
			- перебираем полученный срез;
			- если день недели <= 31 (из расчёта MAX 31 день в месяце) и день месяца >= дню предыдущего месяца,
			добавляем подходящие события в итоговый срез;
			-если день > weekCount и день <= дню, заданному пользователем, добавляем эти результаты в итоговой срез.
	3.3. Переводим результат в JSON и выдаём оба: срез и JSON.
*/

func ShowEventByWeek(date CalendarModel, CalendarStorage []*CalendarModel) (result []CalendarModel, inJSON []byte) {

	var weekCount int
	var otherDays []CalendarModel
	var otherDaysInPrevMonth []CalendarModel

	for _, event := range CalendarStorage {
		if (event.Day == date.Day) && (event.Month == date.Month) && (event.Year == date.Year) {
			weekCount = event.Day - 6
		}

	}

	if (weekCount > 7) || (weekCount == 0) {

		otherDays, _ = ShowEventByMonth(date, CalendarStorage)

		for _, week := range otherDays {
			if (week.Day > weekCount) && (week.Day <= date.Day) {
				result = append(result, week)
			}
		}
	}

	if (weekCount < 7) && (weekCount != 0) {

		prevMonth := CalendarModel{
			Description: "test",
			Day:         weekCount + 30,
			Month:       date.Month - 1,
			Year:        date.Year,
			DayOfWeek:   time.Weekday(date.Day),
		}

		otherDaysInPrevMonth, _ = ShowEventByMonth(prevMonth, CalendarStorage)

		for _, week := range otherDaysInPrevMonth {
			if (week.Day <= 31) && (week.Day >= prevMonth.Day) {
				result = append(result, week)
			}
		}
		otherDays, _ = ShowEventByMonth(date, CalendarStorage)

		for _, week := range otherDays {
			if (week.Day > weekCount) && (week.Day <= date.Day) {
				result = append(result, week)
			}
		}
	}

	inJSON, _ = json.Marshal(result)

	return result, inJSON
}
