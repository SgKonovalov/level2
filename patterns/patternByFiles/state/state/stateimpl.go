package state

import "fmt"

type DriveState struct{}

func (ds DriveState) Action() {
	fmt.Println("Drive ahead")
}

type StopState struct{}

func (sts StopState) Action() {
	fmt.Println("Stop")
}

type ReverseState struct{}

func (rs ReverseState) Action() {
	fmt.Println("Drive reverse")
}

type TurnLeftState struct{}

func (tls TurnLeftState) Action() {
	fmt.Println("Turn left")
}

type TurnRightState struct{}

func (trs TurnRightState) Action() {
	fmt.Println("Turn right")
}
