package main

/*
Паттерн «Цепочка вызовов», является поведенческим, т.е. решает задачи эффективного и
безопасного взаимодействия между объектами программы. Он позволяет:
1) Передавать запросы последовательно по цепочке обработчиков. При этом, каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи
2) Превратить отдельные поведения в объекты;
3) Гарантировать, что передав запрос в первый обработчик цепочки, все объекты в цепи смогут его обработать. При этом длина цепочки не имеет никакого значения.

Пример: процесс сборки автомобиля, состоит из:
1) Сварка кузова;
2) Установки двигателя;
3) Монтажа салона;
4) Установки колёс.

Для реализации:
1) Определяем «сборочную линию», которая запускает процесс и переходит на следующий этап;
2) Определяем сам автомобиль, как структуру, содержащую bool (информацию о прохождении участка);
3) Создаём сами участки линии, реализующие «сборочную линию»

«Плюсы использования»:
1) Нет зависимости от объекта и обработчика;
2) Обработка по установленному алгоритму;

«Минусы использования»:
1) Запрос может быть не обработан.
*/

import (
	"chain/assemblline"
	"chain/definition"
)

func main() {

	car := &definition.CarCR{
		Model: "new",
	}

	bodyShop := &assemblline.BodyShopCR{}
	engineInstallation := &assemblline.EngineInstallationCR{}
	interior := &assemblline.InteriorCR{}
	wheelsCR := &assemblline.WheelsCR{}

	bodyShop.SetNext(engineInstallation)
	engineInstallation.SetNext(interior)
	interior.SetNext(wheelsCR)

	bodyShop.Execute(car)

}