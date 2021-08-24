/*
NTP (сетевой протокол времени) - это протокол синхронизации времени.

Цель использования NTP - синхронизировать часы всех устройств с часами в сети,
чтобы согласовать часы всех устройств в сети, чтобы устройства могли предоставлять
несколько приложений на основе единого времени.

Узнать время можно 2-мя вариантами:

1) «Простой», если нам необходимо ТОЛЬКО узнать текущее время.
Функция getCurrentTimeVarOne() – не принимает аргументов, на «выходе» получаем: время типа time.Time
и ошибку типа err error.
Процесс:
	1. Узнаём время на сервере с помощью библиотеки NTP, вызывая метод ntp.Time, как параметр передавая HOST сервера;
	2. Если такой сервер не найден – выходим из функции, возвращая «пустой» ntpTime и ошибку; в противном случае, возвращаем актуальный ntpTime и nil в качестве ошибки.

2) Для получения не только актуального времени, но и дополнительных метаданных.
Функция getCurrentTimeVarTwo():
	1. Узнаём время на сервере с помощью библиотеки NTP, вызывая метод ntp.Query, как параметр передавая HOST сервера;
	2. Если такой сервер не найден – выходим из функции, возвращая «пустой» ntpTime и ошибку; в противном случае, переходим к следующему шагу;
	3. Присваиваем ntpTime значение текущего времени с добавлением Duration.
*/

package main

import (
	"fmt"
	"log"
	"time"

	ntp "github.com/beevik/ntp"
)

func getCurrentTimeVarOne() (ntpTime time.Time, err error) {

	ntpTime, err = ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return ntpTime, err
	}
	return ntpTime, nil
}

func getCurrentTimeVarTwo() (ntpTime time.Time, err error) {

	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return ntpTime, err
	}

	ntpTime = time.Now().Add(response.ClockOffset)

	return ntpTime, nil
}

func main() {

	timeVarOne, errVarOne := getCurrentTimeVarOne()

	if errVarOne != nil {
		log.Fatal(errVarOne)
	}
	fmt.Println(timeVarOne)

	timeVarTwo, errVarTwo := getCurrentTimeVarTwo()
	if errVarTwo != nil {
		log.Fatal(errVarTwo)
	}
	fmt.Println(timeVarTwo)
}
