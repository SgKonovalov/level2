/*
Паттерн «Состояние», является поведенческим паттерном,
т.е. решает задачи эффективного и безопасного взаимодействия между объектами программы.
Он позволяет:
1) Объектам менять поведение в зависимости от своего состояния;
2) Реализация раскладывается на множество классов-состояний.

Пример: у автомобиля может быть 5 состояний:
1) Ехать вперёд;
2) Остановиться;
3) Ехать назад;
4) Поернуть налево;
5) Повернуть направо.
Так, все структуры state - характеризуют множество его состояний,
а context - устанавливает текущее состояние автомобиля (едет вперёд или остановился).

«Плюсы использования»:
1) Упрощает код контекста.

«Минусы использования»:
1) Код усложняется.
*/

package main

import (
	"fmt"
)

type carState interface {
	action()
}

type driveState struct{}

func (ds driveState) action() {
	fmt.Println("Drive ahead")
}

type stopState struct{}

func (sts stopState) action() {
	fmt.Println("Stop")
}

type reverseState struct{}

func (rs reverseState) action() {
	fmt.Println("Drive reverse")
}

type turnLeftState struct{}

func (tls turnLeftState) action() {
	fmt.Println("Turn left")
}

type turnRightState struct{}

func (trs turnRightState) action() {
	fmt.Println("Turn right")
}

type carMovingContext struct {
	cs carState
}

func (cmc *carMovingContext) setState(cs carState) {
	cmc.cs = cs
}

func (cmc carMovingContext) action() {
	cmc.cs.action()
}

/*func main() {
	csm := carMovingContext{}

	carD := driveState{}
	carS := stopState{}
	carR := reverseState{}
	tLeft := turnLeftState{}
	tRight := turnRightState{}

	csm.setState(carD)
	csm.action()

	csm.setState(carS)
	csm.action()

	csm.setState(carR)
	csm.action()

	csm.setState(tLeft)
	csm.action()

	csm.setState(tRight)
	csm.action()
}*/
