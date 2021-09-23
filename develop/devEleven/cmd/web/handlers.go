package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"exeTwo.devEleven/pkg"
)

/*
home - стандартный handler обработки стартовой страницы.
*/

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

}

/*
events_for_day - handler обработки запроса на просмотр событий по дням.
Вызывает соответсвующи метод и производит логирование.
*/

func events_for_day(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")

	log.Printf("SEARCH: looking for event by day %s", date)

	if date == "" {
		http.Error(w, "Incorrect date format", 400)
		return
	}

	gettingDate := pkg.NewCalendarModel(date, "test", rand.Int())
	if gettingDate.Day == 0 {
		http.Error(w, "Incorrect date format", 400)
		return
	}

	_, inJSON := pkg.ShowEventByDay(*gettingDate, pkg.CalendarStorage)

	if len(inJSON) <= 4 {
		resErr := pkg.ResultError{
			Result: "No events for choosed day",
		}
		errInJson, _ := json.Marshal(resErr)
		fmt.Fprintln(w, string(errInJson))
		return
	}

	resOk := pkg.ResultOk{
		Result: string(inJSON),
	}

	okInJSON, _ := json.Marshal(resOk)

	log.Println("SEARCH: looking for event by day executing SUCCESSFUL")

	fmt.Fprintln(w, string(okInJSON))
}

/*
events_for_week - handler обработки запроса на просмотр событий по неделе.
Вызывает соответсвующи метод и производит логирование.
*/

func events_for_week(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")

	log.Printf("SEARCH: looking for event by week %s", date)

	if date == "" {
		http.Error(w, "Incorrect date format", 400)
		return
	}

	gettingDate := pkg.NewCalendarModel(date, "test", rand.Int())
	if gettingDate.Day == 0 {
		http.Error(w, "Incorrect date format", 400)
		return
	}

	_, inJSON := pkg.ShowEventByWeek(*gettingDate, pkg.CalendarStorage)

	if len(inJSON) <= 4 {
		resErr := pkg.ResultError{
			Result: "No events for choosed week",
		}
		errInJson, _ := json.Marshal(resErr)
		fmt.Fprintln(w, string(errInJson))
		return
	}

	resOk := pkg.ResultOk{
		Result: string(inJSON),
	}

	okInJSON, _ := json.Marshal(resOk)

	log.Println("SEARCH: looking for event by week executing SUCCESSFUL")

	fmt.Fprintln(w, string(okInJSON))
}

/*
events_for_month - handler обработки запроса на просмотр событий по месяцу.
Вызывает соответсвующи метод и производит логирование.
*/

func events_for_month(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")

	log.Printf("SEARCH: looking for event by month %s", date)

	if date == "" {
		http.Error(w, "Incorrect date format", 400)
		return
	}

	gettingDate := pkg.NewCalendarModel(date, "test", rand.Int())
	if gettingDate.Day == 0 {
		http.Error(w, "Incorrect date format", 400)
		return
	}

	_, inJSON := pkg.ShowEventByMonth(*gettingDate, pkg.CalendarStorage)

	if len(inJSON) <= 4 {
		resErr := pkg.ResultError{
			Result: "No events for choosed month",
		}
		errInJson, _ := json.Marshal(resErr)
		fmt.Fprintln(w, string(errInJson))
		return
	}

	resOk := pkg.ResultOk{
		Result: string(inJSON),
	}

	okInJSON, _ := json.Marshal(resOk)

	log.Println("SEARCH: looking for event by month executing SUCCESSFUL")

	fmt.Fprintln(w, string(okInJSON))
}

/*
create_event - handler обработки запроса на добавление новых событий.
Вызывает соответсвующи метод и производит логирование.
*/

func create_event(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed!", 500)
		return
	}

	log.Println("CRUD: create function in process")

	decoder := json.NewDecoder(r.Body)
	var newEvent *pkg.CalendarModel
	err := decoder.Decode(&newEvent)
	if err != nil {
		http.Error(w, "Wrong JSON format", 400)
	}

	err = pkg.CreateNewEvent(newEvent)
	if err != nil {
		resErr := pkg.ResultError{
			Result: string(err.Error()),
		}
		errInJson, _ := json.Marshal(resErr)
		fmt.Fprintln(w, string(errInJson))
		http.Error(w, string(err.Error()), 503)
		return
	}

	resOk := pkg.ResultOk{
		Result: "Event added succesfully",
	}

	okInJSON, _ := json.Marshal(resOk)

	log.Println("CRUD: create executed SUCCESSFUL")

	fmt.Fprintln(w, string(okInJSON))
}

/*
update_event - handler обработки запроса на добавление изменение событий.
Вызывает соответсвующи метод и производит логирование.
*/

func update_event(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed!", 500)
		return
	}

	log.Println("CRUD: update function in process")

	decoder := json.NewDecoder(r.Body)
	var newEvent *pkg.CalendarModel
	err := decoder.Decode(&newEvent)
	if err != nil {
		http.Error(w, "Wrong JSON format", 400)
	}

	err = pkg.UpdateEvent(newEvent)
	if err != nil {
		resErr := pkg.ResultError{
			Result: string(err.Error()),
		}
		errInJson, _ := json.Marshal(resErr)
		fmt.Fprintln(w, string(errInJson))
		http.Error(w, string(err.Error()), 503)
		return
	}

	resOk := pkg.ResultOk{
		Result: "Event updated succesfully",
	}

	okInJSON, _ := json.Marshal(resOk)

	log.Println("CRUD: update executed SUCCESSFUL")

	fmt.Fprintln(w, string(okInJSON))
}

/*
delete_event - handler обработки запроса на добавление удаление событий.
Вызывает соответсвующи метод и производит логирование.
*/

func delete_event(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed!", 500)
		return
	}

	log.Println("CRUD: delete function in process")

	decoder := json.NewDecoder(r.Body)
	var newEvent *pkg.CalendarModel
	err := decoder.Decode(&newEvent)
	if err != nil {
		http.Error(w, "Wrong JSON format", 400)
	}

	err = pkg.DeleteEvent(newEvent)
	if err != nil {
		resErr := pkg.ResultError{
			Result: string(err.Error()),
		}
		errInJson, _ := json.Marshal(resErr)
		fmt.Fprintln(w, string(errInJson))
		http.Error(w, string(err.Error()), 503)
		return
	}

	resOk := pkg.ResultOk{
		Result: "Event deleted succesfully",
	}

	okInJSON, _ := json.Marshal(resOk)

	log.Println("CRUD: delete executed SUCCESSFUL")

	fmt.Fprintln(w, string(okInJSON))
}
