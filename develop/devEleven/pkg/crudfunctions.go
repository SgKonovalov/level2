package pkg

import (
	"errors"
)

/*
1) Функция CreateNewEvent() – принимая в качестве аргумента указатель на объект типа CalendarModel добавляет
в Хранилище данные о новом событии:
	1.1. Проверяет ID этого события в Хранилище и если такой ID уже есть в базе,
	выходит из функции, выдавая ошибку: «event with same ID already exist».
	1.2. При помощи append() добавляет новое событие в хранилище.
*/

func CreateNewEvent(date *CalendarModel) (err error) {

	for _, existEvent := range CalendarStorage {
		if date.EventID == existEvent.EventID {
			return errors.New("event with same ID already exist")
		}
	}

	CalendarStorage = append(CalendarStorage, date)

	return err
}

/*
2) Функция DeleteEvent() – принимая в качестве аргумента указатель на объект типа CalendarModel,
удаляет его из Хранилища:
	2.1. Сохраняем в переменную arrLenBefore длину Хранилища ДО удаления;
	2.2. Перебираем Хранилище, если указанный пользователем ID события для удаления встречается в Хранилище:
		2.2.1. Элементу, совпавшему с указанным пользователем, присваиваем номер:
		номер этого элемента «минус» один (т.е. перемещаем его в конец среза);
		2.2.2. Этому номер присваиваем значение nil;
		2.2.3. Копируем в исходное хранилище все элементы до последнего (из п. 1.2.2)
	2.3. Сравниваем длину полученного в п. 1.2.3 среза и исходную длину Хранилища (переменная arrLenBefore)
		2.3.1. Если они не равны, выходим из функции, присваивая ошибке значение nil;
		2.3.2. Если они равны, выдаём ошибку: «event with same ID not exist».
*/

func DeleteEvent(date *CalendarModel) (err error) {

	arrLenBefore := len(CalendarStorage)

	for elemNumber, findToDelete := range CalendarStorage {

		if findToDelete.EventID == date.EventID {
			CalendarStorage[elemNumber] = CalendarStorage[len(CalendarStorage)-1]
			CalendarStorage[len(CalendarStorage)-1] = nil
			CalendarStorage = CalendarStorage[:len(CalendarStorage)-1]
		}
	}

	arrLenAfter := len(CalendarStorage)

	if arrLenBefore != arrLenAfter {
		return nil
	} else {
		return errors.New("event with same ID not exist")
	}
}

/*
3) Функция UpdateEvent() – принимая в качестве аргумента указатель на объект типа CalendarModel,
изменяет его в Хранилище:
	3.1. Перебираем Хранилище. Если находим событие, заданное пользователем,
	добавляем его значение в промежуточный срез elemBeforeChange, а так же присваиваем этому событие значение,
	заданное пользователем.
	3.2. Если длина промежуточного среза == 0 => элемента, заданного пользователем в Хранилище нет =>
	выдаём ошибку: «no events with entered ID»;
	3.3. Перебираем промежуточный массив.
	Если все значения его элемента совпадают с указанными пользователем параметрами =>
	никаких изменений в Хранилище внесено не было => выдаём ошибку: «event not changed».
*/

func UpdateEvent(date *CalendarModel) (err error) {

	var elemBeforeChange []CalendarModel

	for elemNumb, findToUpdate := range CalendarStorage {

		if findToUpdate.EventID == date.EventID {
			elemBeforeChange = append(elemBeforeChange, *CalendarStorage[elemNumb])
			CalendarStorage[elemNumb] = date
		}
	}

	if len(elemBeforeChange) == 0 {
		return errors.New("no events with entered ID")
	}

	for _, compare := range elemBeforeChange {
		if (compare.Day == date.Day) && (compare.DayOfWeek == date.DayOfWeek) && (compare.Month == date.Month) && (compare.Year == date.Year) && (compare.Description == date.Description) {
			return errors.New("event not changed")
		}
	}

	return nil
}
