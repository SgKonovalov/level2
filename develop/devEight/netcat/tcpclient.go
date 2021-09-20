package netcat

import (
	"fmt"
	"net"
)

/*
Функция StartTCPClient():
1) Присоединяется к серверу tcp;
2) В бесконечном цикле предаёт сообщения, набранные пользователем в консоли:
	2.1. Сканирует сообщение пользователя fmt.Scanln;
	2.2. Отправляет сообщения пользователя на сервер: conn.Write.
*/

func StartTCPClient() {

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
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
