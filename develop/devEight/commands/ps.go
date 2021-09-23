package commands

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
)

/*
Функция PS():
1) Открывает PowerShell, при помощи exec.Command, используя cmd.StdinPipe();
2) Выполняет команду Get-Process;
3) Выводит в консоль все текущие процессы.
*/

func PS() (err error) {

	cmd := exec.Command("powershell", "-nologo", "-noprofile")

	pipe, err := cmd.StdinPipe()
	if err != nil {
		return errors.New(cantOpenedPowerShell)
	}
	pipe.Write([]byte("Get-Process"))
	pipe.Close()

	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(output))

	return nil
}
