package main

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

import (
	"command/command"
	"command/definition"
)

func main() {

	engineOne := &definition.EngineOneC{}
	engineTwo := &definition.EngineTwoC{}

	startEngineOne := &command.StartEngineC{
		EngineC: engineOne,
	}
	startEngineTwo := &command.StartEngineC{
		EngineC: engineTwo,
	}

	stopEngineOne := &command.StopEngineC{
		EngineC: engineOne,
	}

	stopEngineTwo := &command.StopEngineC{
		EngineC: engineTwo,
	}

	startEngineOne.Execute()
	startEngineTwo.Execute()
	stopEngineOne.Execute()
	stopEngineTwo.Execute()
}
