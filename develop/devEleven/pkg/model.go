package pkg

import (
	"time"
)

/*
CalendarStorage - имитация хранилища событий. Далее - Хранилище.
*/

var CalendarStorage = []*CalendarModel{
	NewCalendarModel("2021-07-08", "aug1", 16),
	NewCalendarModel("2021-07-08", "aug1", 23),
}

/*
CalendarModel - Структура самого события для хранения.
*/

type CalendarModel struct {
	EventID     int          `json:"id"`
	Description string       `json:"description"`
	Day         int          `json:"day"`
	Month       time.Month   `json:"month"`
	Year        int          `json:"year"`
	DayOfWeek   time.Weekday `json:"dayofweek"`
}

/*
NewCalendarModel - Конструктор нового события.
*/

func NewCalendarModel(date, Description string, ID int) *CalendarModel {

	retCM := new(CalendarModel)

	dateForParse, err := time.Parse("2006-02-01", date)
	if err != nil {
		return retCM
	}
	retCM.EventID = ID
	retCM.Description = Description
	retCM.Day = dateForParse.Day()
	retCM.Month = dateForParse.Month()
	retCM.Year = dateForParse.Year()
	retCM.DayOfWeek = dateForParse.Weekday()

	return retCM
}

/*
ResultOk - Структура для возвращения JSON документа, в случае успешного выполнения функции.
*/

type ResultOk struct {
	Result string `json:"result"`
}

/*
ResultError - Структура для возвращения JSON документа, в случае ошибки бизнес-логики.
*/

type ResultError struct {
	Result string `json:"error"`
}
