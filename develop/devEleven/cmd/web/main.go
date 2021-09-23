package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

/*
Функция main():
1) Устанавливает адрес сервера;
2) Отвечает за основной лог работы программы;
3) Запускает среверы и отвечает за его работу;
4) обслуживает routes.
*/

func main() {
	addr := flag.String("addr", ":8080", "Servers address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/events_for_day", events_for_day)
	mux.HandleFunc("/events_for_week", events_for_week)
	mux.HandleFunc("/events_for_month", events_for_month)
	mux.HandleFunc("/create_event", create_event)
	mux.HandleFunc("/update_event", update_event)
	mux.HandleFunc("/delete_event", delete_event)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Start service at %s", *addr)

	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
