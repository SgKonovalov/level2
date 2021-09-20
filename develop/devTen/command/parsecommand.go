package command

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

/*
1) Функция ReadCommandFromConsole() – считывает команду, введённую пользователем в консоли.
*/

func ReadCommandFromConsole() (command string) {

	fmt.Println("Insert your command and then press 'Enter', please:")

	myscan := bufio.NewScanner(os.Stdin)
	myscan.Scan()
	command = myscan.Text()
	return command
}

/*
1) Функция ParsingGettingCommand() –  получая в качестве аргумента команду, введённую пользователем,
расшифровывает её и в результате выдаёт информацию о нужном host и port для подключения к серверу,
а так же сведения о наличии timeout и о его длительности, указанной пользователем.
	1.1. При помощи регулярных выражений, разделяет команд на 2 группы:
		1.1.1. regexp.MustCompile(`\D\s\d+$`) – если указан одновременно host и port для подключения;
		1.1.2. regexp.MustCompile(`\d+.\d+\s\d+`) – если указан только port для подключения;
	1.2. Вызывает функцию LookForTimeout() – для уточнения наличия сведения о timeout в команде пользователя.
*/

func ParsingGettingCommand(command string) (host, port string, hasTimeout bool, timeout int) {

	isThisHostAndPort := regexp.MustCompile(`\D\s\d+$`)
	hostAndPortToString := isThisHostAndPort.FindString(command)

	if hostAndPortToString != "" {
		host, port = HostAndPort(command)
	}

	isThisOnlyHost := regexp.MustCompile(`\d+.\d+\s\d+`)
	onlyHostToString := isThisOnlyHost.FindString(command)

	if onlyHostToString != "" {
		port = OnlyPort(command)
	}

	hasTimeout, timeout = LookForTimeout(command)

	return host, port, hasTimeout, timeout
}
