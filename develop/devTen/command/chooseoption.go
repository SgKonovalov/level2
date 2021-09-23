package command

import (
	"fmt"

	"exeTwo.devTen/tcpconnection"
)

/*
1) Функция ChooseOption() – получая в качестве аргументов host, port, информацию о наличии timeout в команде и само время timeout, координирует работу функций подключения к серверам, вызывая нужный функции из пакета tcpconnection.
	1.1. Проверяет полученные значения host/port:
		1.1.1. Если host и port – не пустые => присоединяемся к серверу по этим данным;
		1.1.2. Если пустые, то используем только port.
	1.2. В зависимости от налчия опции timeiut:
		1.2.1. Подключается к сервееру с установленным пользователем timeout, если hasTimeout == true и timeout < 0. Если timeout < 0, то подключается к серверву со значением timeout по умолчанию (10 секунд).
		1.2.2. Если hasTimeout == false, подключается к серверу со значением timeout по умолчанию.
*/

func ChooseOption(host, port string, hasTimeout bool, timeout int) {

	var conectTo string

	if host != "" && port != "" {
		conectTo = fmt.Sprint(host + ":" + port)
	} else {
		conectTo = port
	}

	if hasTimeout && timeout != 0 {

		tcpconnection.TCPClient(conectTo, timeout)
	}

	if hasTimeout && timeout == 0 {

		tcpconnection.TCPClient(conectTo, 10)
	}

	if !hasTimeout {

		tcpconnection.TCPClient(conectTo, 10)
	}

}
