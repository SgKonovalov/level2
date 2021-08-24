/*
Паттерн «Команда», является поведенческим, т.е. решает задачи эффективного и
безопасного взаимодействия между объектами программы.
Т.е. «Команда» используется для создания механизма действий над чем-либо.
В этом механизме, структура-отправитель команды и структура-получатель не зависят друг от друга.
Для этого, «Команда»:
1) Инкапсулирует запрос как объект, позволяя тем самым задавать параметры клиентов
для обработки соответствующих запросов, ставить запросы в очередь или протоколировать их,
а также поддерживать отмену операций;
2) Позволяет библиотечным объектам отправлять запросы не известным объектам приложения,
преобразовав сам запрос в объект. Этот объект можно хранить и передавать, как и любой другой;
3) Превращает операции в объекты. А объекты можно передавать, хранить и менять внутри других объектов.

Пример: можно при помощи одного ключа завести 2 двигателя автомобиля:
1) Определяем интерфейс самой команды;
2) Создаём конкретные команды: завести и заглушить двигатель, реализуя интерфейс из п. 1;
3) Создаём интерфейс самого двигателя, который содержит указанные в п. 2 методы;
4) Создаём структуры, реализующие интерфейс из п. 2;
5) В main – создаём новый объекты структур из п. 4 + у ним, структуры из п. 2.

«Плюсы использования»:
1) Объекты не зависят друг от друга;
2) Объекты могут «общаться» между собой;

«Минусы использования»:
1) Сложный код программы.
*/

package main

import "fmt"

type command interface {
	execute()
}

type startEngineC struct {
	engineC engineC
}

func (sec *startEngineC) execute() {
	sec.engineC.startEngineC()
}

type stopEngineC struct {
	engineC engineC
}

func (stec *stopEngineC) execute() {
	stec.engineC.stopEngineC()
}

type engineC interface {
	startEngineC()
	stopEngineC()
}

type engineOneC struct {
	isStarting bool
}

type engineTwoC struct {
	isStarting bool
}

func (ec *engineOneC) startEngineC() {
	ec.isStarting = true
	fmt.Println("Engine one been started")
}

func (ec *engineOneC) stopEngineC() {
	ec.isStarting = false
	fmt.Println("Engine one been stopped")
}

func (ec *engineTwoC) startEngineC() {
	ec.isStarting = true
	fmt.Println("Engine two been started")
}

func (ec *engineTwoC) stopEngineC() {
	ec.isStarting = false
	fmt.Println("Engine two been stopped")
}

/*func main() {

	engineOne := &engineOneC{}
	engineTwo := &engineTwoC{}

	startEngineOne := &startEngineC{
		engineC: engineOne,
	}
	startEngineTwo := &startEngineC{
		engineC: engineTwo,
	}

	stopEngineOne := &stopEngineC{
		engineC: engineOne,
	}

	stopEngineTwo := &stopEngineC{
		engineC: engineTwo,
	}

	startEngineOne.execute()
	startEngineTwo.execute()
	stopEngineOne.execute()
	stopEngineTwo.execute()
}*/
