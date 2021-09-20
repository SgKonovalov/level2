package main

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

import (
	context "state/context"
	"state/state"
)

func main() {
	csm := context.CarMovingContext{}

	carD := state.DriveState{}
	carS := state.StopState{}
	carR := state.ReverseState{}
	tLeft := state.TurnLeftState{}
	tRight := state.TurnRightState{}

	csm.SetState(carD)
	csm.Action()

	csm.SetState(carS)
	csm.Action()

	csm.SetState(carR)
	csm.Action()

	csm.SetState(tLeft)
	csm.Action()

	csm.SetState(tRight)
	csm.Action()
}
