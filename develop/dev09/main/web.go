package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/*
1) Функция insertIntoFile() – получая в качестве аргумента html-код страницы,
создаёт html-файл и записывает в него код станицы для дальнейшей работы с ней:
	1.1. os.Create("./static/site.html") – создаёт файл в установленной директории программы;
	1.2. file.WriteString(htmlCode) – записывает в этот файл html-код страницы.
*/

func insertIntoFile(htmlCode string) {

	file, err := os.Create("./static/site.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(htmlCode)
	if err != nil {
		panic(err)
	}
}

/*
2) Функция Home() – обчный handler с route "/". При переходе на него:
	2.1. Запрашивается нужная страница, при помощи функции http.Get();
	2.2. При помощи ioutil.ReadAll(req.Body) считывает в переменную site html-код страницы;
	2.3. Вызывает функцию insertIntoFile(), которой передам переменную site.
*/

func Home(w http.ResponseWriter, r *http.Request) {

	req, err := http.Get("https://docs.google.com/document/d/1R7dbON52NlBLYun8crm4nHtedS1wZ9TCdQrWRLjIEio/edit")
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	site, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(site)

	insertIntoFile(string(site))
}

/*
4) Функция Server() – координирует работу сервера, пишет logs и прослушивает порт.
*/

func Server(portNum string, Home func(w http.ResponseWriter, r *http.Request)) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	log.Printf("Запуск веб-сервера на %s", portNum)
	err := http.ListenAndServe(portNum, mux)
	log.Fatal(err)
}
