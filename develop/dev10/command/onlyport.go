package command

import "regexp"

/*
1) Функция OnlyPort() – получая в качестве аргумента команду, введённую пользователем,
расшифровывает её и в результате выдаёт информацию о нужном port для подключения к серверу.
	1.1. При помощи регулярного выражения regexp.MustCompile(`\d+.\d+.\d+.\d+\s\d+`),
	ищем в строке команды port для подключения.
	1.2. Возвращаем полученный port.
*/

func OnlyPort(command string) (port string) {

	lookForPort := regexp.MustCompile(`\d+.\d+.\d+.\d+\s\d+`)
	port = lookForPort.FindString(command)

	return port
}
