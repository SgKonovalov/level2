package main

import (
	"testing"

	"exeTwo.devEight/commands"
)

func TestSortByPWDKey(t *testing.T) {
	var check = `D:\devEight`

	letTest, err := commands.PWD()
	if err != nil {
		t.Fatalf("ERROR! %v", err)
	}

	if letTest != check {
		t.Fatalf("Incorrect parsing %v. Expected: %v", letTest, check)
	}
}
