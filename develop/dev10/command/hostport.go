package command

import (
	"regexp"
)

/*
1) Функция HostAndPort() – получая в качестве аргумента команду, введённую пользователем,
расшифровывает её и в результате выдаёт информацию о нужном host и port для подключения к серверу.
	1.1. При помощи регулярный выражений, ищем значения:
		1.1.1. Ищем номер порта при помощи regexp.MustCompile(`\d+$`);
		1.1.2. Ищем в команде DNS адрес – regexp.MustCompile(`\w+\.\w+`);
		1.1.3. Если не находит host по DNS то ищем его при помощи регулярного выражения:
		regexp.MustCompile(`\w+host`).
	1.2. Возвращаем полученный в п. 1.1.1 и в п. 1.1.2 или 1.1.3.
*/

func HostAndPort(command string) (host, port string) {

	lookForPort := regexp.MustCompile(`\d+$`)
	port = lookForPort.FindString(command)

	lookForHostDNS := regexp.MustCompile(`\w+\.\w+`)
	host = lookForHostDNS.FindString(command)

	if host == "" {
		lookForHostOnlyHost := regexp.MustCompile(`\w+host`)
		host = lookForHostOnlyHost.FindString(command)
	}

	return host, port
}
