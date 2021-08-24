package netcat

import (
	"fmt"
	"log"
	"net"
)

/*
Функция StartUDPClient():
1) Присоединяется к порту;
2) Присоединяется к удалённому серверу UDP;
3) В бесконечном цикле предаёт сообщения, набранные пользователем в консоли:
    3.1. Сканирует сообщение пользователя fmt.Scanln;
    3.2. Отправляет сообщения пользователя на сервер: conn.Write.
*/

func StartUDPClient() {

	service := "localhost:8080"
	RemoteAddr, _ := net.ResolveUDPAddr("udp", service)

	conn, err := net.DialUDP("udp", nil, RemoteAddr)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	for {
		var source string
		fmt.Print("Insert yor message: ")
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println("Incerrect message", err)
			continue
		}

		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
	}
}
