package tcpconnection

import (
	"fmt"
	"net"
	"os"
	"time"
)

/*
Функция TCPClient():
1) Присоединяется к серверу tcp, с установленным timeout;
2) В бесконечном цикле предаёт сообщения, набранные пользователем в консоли:
	2.1. Сканирует сообщение пользователя fmt.Scanln;
	2.2. Отправляет сообщения пользователя на сервер: conn.Write.
	2.3. Читает сообщения, полученный от сервера: conn.Read
*/

func TCPClient(conectTo string, timeout int) {

	fmt.Println(conectTo)

	conn, err := net.DialTimeout("tcp", conectTo, time.Duration(timeout)*time.Second)

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

		exit := []rune(source)

		for _, foundD := range exit {
			if foundD == 4 {
				os.Exit(1)
			}
		}

		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}

		reply := make([]byte, 1024)

		_, err = conn.Read(reply)
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		println("reply from server=", string(reply))

	}

}
