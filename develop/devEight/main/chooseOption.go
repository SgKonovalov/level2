package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"exeTwo.devEight/commands"
	"exeTwo.devEight/netcat"
)

/*
1) Функция ReadCommandFromConsole() - читает выбранную пользователем команду. Выдаёт результат типа string.
*/

func ReadCommandFromConsole() (command string) {

	fmt.Println("Insert your command and then press 'Enter', please:")

	myscan := bufio.NewScanner(os.Stdin)
	myscan.Scan()
	command = myscan.Text()
	return command
}

/*
2) Функция ReadProtocolFromConsole() - читает выбранный протокол для пердачи данных(TCP или UDP).
*/

func ReadProtocolFromConsole() (protocol string) {

	fmt.Println("Insert protocol, which you want to use: 'udp' for UDP, 'tcp' for TCP and then press 'Enter', please:")

	myscan := bufio.NewScanner(os.Stdin)
	myscan.Scan()
	protocol = myscan.Text()
	return protocol
}

/*
3) Функция ncCommandParsing() – для утилиты netcat.
Расшифроввает саму команду и выдает 2 значения типа bool для протоколов TCP и UDP.
В качестве аргумента принимает строку, сформированную ReadProtocolFromConsole().
*/

func ncCommandParsing(protocol string) (udp, tcp bool) {

	if strings.Contains(protocol, "udp") {
		udp = true
	}

	if strings.Contains(protocol, "tcp") {
		tcp = true
	}

	return udp, tcp
}

/*
4) Функция ReadProcessIDFromConsole() – читает и расшифровывает ID процесса для дальнейшей остановки.
*/

func ReadProcessIDFromConsole() (processID int) {

	fmt.Println("Insert ID of procces, which you want to STOP, like: '12345' and then press 'Enter', please:")

	myscan := bufio.NewScanner(os.Stdin)
	myscan.Scan()
	processID, err := strconv.Atoi(myscan.Text())
	if err != nil {
		log.Fatal(err)
	}
	return processID
}

/*
5) Функция ReadPathFromConsole() – читает и расшифровывает путь к папке,
на которую пользователь хочет заменить текущую директорию.
*/

func ReadPathFromConsole() (path string) {
	fmt.Println("Insert path to file/folder (path to folder for 'cd' command, like: 'C:/Google/Chrome'\n or path to file for 'exec', like: 'C:/Google/Chrome/run.exe') and then press 'Enter', please:")

	myscan := bufio.NewScanner(os.Stdin)
	myscan.Scan()
	path = myscan.Text()
	return path
}

/*
6) Функция ParsingGettingCommand() – расшифровывает полученную команду и присваивает значения типа bool для каждой.
Далее эти значения используются в координаторе chooseCommand().
*/

func ParsingGettingCommand(command string) (cd, pwd, echo, kill, ps, exec, fork, nc bool) {

	if strings.Contains(command, "cd") {
		cd = true
	}

	if strings.Contains(command, "pwd") {
		pwd = true
	}

	if strings.Contains(command, "echo") {
		echo = true
	}

	if strings.Contains(command, "kill") {
		kill = true
	}

	if strings.Contains(command, "ps") {
		ps = true
	}

	if strings.Contains(command, "exec") {
		exec = true
	}

	if strings.Contains(command, "fork") {
		fork = true
	}

	if strings.Contains(command, "nc") {
		nc = true
	}

	return cd, pwd, echo, kill, ps, exec, fork, nc
}

/*
7) Функция chooseCommand() – координирует работу с командами и вызывает нужную функцию пакета command,
в соответствии с введёнными ключами.
*/

func chooseCommand(cd, pwd, echo, kill, ps, exec, fork, nc bool) (result string, err error) {

	if cd {
		result, err = commands.CD(ReadPathFromConsole())
		if err != nil {
			return "", err
		}
		return result, nil
	}

	if exec {
		err = commands.EXEC(ReadPathFromConsole())
		if err != nil {
			return "", err
		}
		return "", nil
	}
	if pwd {
		result, err = commands.PWD()
		if err != nil {
			return "", err
		}
		return result, nil
	}

	if echo {
		commands.ECHO()
		return "", nil
	}

	if fork {
		err = commands.FORK()
		if err != nil {
			return "", nil
		}
		return "", nil
	}

	if kill {
		err = commands.KILL(ReadProcessIDFromConsole())
		if err != nil {
			return "", err
		}
		return "", nil
	}

	if ps {
		err = commands.PS()
		if err != nil {
			return "", err
		}
		return "", nil
	}

	if nc {

		udp, tcp := ncCommandParsing(ReadProtocolFromConsole())

		if udp && tcp {
			return "", errors.New(commands.ChoosingTCPandUDP)
		}

		if tcp {
			netcat.StartTCPClient()
			return "", nil
		}

		if udp {
			netcat.StartUDPClient()
			return "", nil
		}
	}

	return result, err
}
