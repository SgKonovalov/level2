package commands

import (
	"errors"
	"fmt"
	"os/exec"
)

/*
Функция KILL() – принимает в качестве аргумента ID процесса, который необходимо завершить.
1) Открывает PowerShell, при помощи exec.Command, используя cmd.StdinPipe();
2) Выполняет команду Stop-Process…, конструируя строку из заготовленного шаблона,
добавляя в него ID полученного от пользователя процесса.
*/

func KILL(processID int) (err error) {

	cmd := exec.Command("powershell", "-nologo", "-noprofile")

	pipe, err := cmd.StdinPipe()
	if err != nil {
		return errors.New(cantOpenedPowerShell)
	}

	command := fmt.Sprintf("Stop-Process %d -force", processID)

	pipe.Write([]byte(command))
	pipe.Close()

	return nil
}
