package main

/*
Паттерн «Стратегия», является поведенческим паттерном, т.е. решает задачи эффективного и безопасного взаимодействия между объектами программы.
Он позволяет выбор поведения алгоритма в ходе исполнения.
1) Определяет семейство алгоритмов, инкапсулирует каждый из них и делает их взаимозаменяемыми. Стратегия позволяет изменять алгоритмы независимо от клиентов, которые ими пользуются;
2) Алгоритмы инкапсулируется классами, которые и являются стратегиями;
3) Классы стратегий реализуют общий интерфейс;
4) Контекст содержит ссылку на конкретную стратегию и вызывает ей при необходимости.

Используем, если:
1) Есть несколько родственных классов, которые отличаются поведением;
2) Нужны несколько вариантов поведения;
3) Необходимо скрыть от клиента данные класса;
4) Используется много условных операторов.

Пример:
Автомобиль может быть окрашен в 5 цветов.
1) Для этого определяем сами структуры цветов:
	1.1. Общее понятие цвета автомобиля (carColour);
	1.2. Красный (redCar);
	1.3. Белый (whiteCar);
	1.4. Чёрный (blackCar);
	1.5. Синий (blueCar);
	1.6. Зелёный (greenCar).
2) Определяем струткуру самого автомобиля (car), с возможностью:
	2.1. Выбора цвета (setState);
	2.2. Самого окрашивания (painting).
3) В клиенте main, определяется порядок окраски.

«Плюсы использования»:
1) Делегирование вместо композиции;
2) Инкапсуляция кода и данных;

«Минусы использования»:
1) Код усложняется.

Отличие от паттерна состояние:
1) Стратегия позволяет клиенту выбрать другое поведение объекта, а состояние - помогает объекту управлять состоянием;
2) В Состоянии переходом управляет контекст, а в Стратегии - клиент.
*/

import (
	"strategy/action"
	"strategy/definition"
)

func main() {
	car := definition.Car{}
	rc := action.RedCar{}
	bc := action.BlackCar{}
	wc := action.WhiteCar{}
	bueC := action.BlueCar{}
	gc := action.GreenCar{}

	car.SetState(rc)
	car.Painting()

	car.SetState(bc)
	car.Painting()

	car.SetState(wc)
	car.Painting()

	car.SetState(bueC)
	car.Painting()

	car.SetState(gc)
	car.Painting()
}
