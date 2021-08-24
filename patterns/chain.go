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

package main

import "fmt"

type assemblingLineCR interface {
	execute(*carCR)
	setNext(assemblingLineCR)
}

type carCR struct {
	model             string
	bodyShopIsDone    bool
	engineInstalled   bool
	interiorInstalled bool
	wheelsInstaled    bool
}

type bodyShopCR struct {
	next assemblingLineCR
}

func (bscr *bodyShopCR) execute(ccr *carCR) {
	if ccr.bodyShopIsDone {
		fmt.Println("Took body from stock")
		bscr.next.execute(ccr)
		return
	}
	fmt.Println("Welding body by robots")
	ccr.bodyShopIsDone = true
	bscr.next.execute(ccr)
}

func (bscr *bodyShopCR) setNext(next assemblingLineCR) {
	bscr.next = next
}

type engineInstallationCR struct {
	next assemblingLineCR
}

func (eicr *engineInstallationCR) execute(ccr *carCR) {
	if ccr.engineInstalled {
		fmt.Println("Engine already in body")
		eicr.next.execute(ccr)
		return
	}
	fmt.Println("Installation engine by workers")
	ccr.engineInstalled = true
	eicr.next.execute(ccr)
}

func (eicr *engineInstallationCR) setNext(next assemblingLineCR) {
	eicr.next = next
}

type interiorCR struct {
	next assemblingLineCR
}

func (icr *interiorCR) execute(ccr *carCR) {
	if ccr.interiorInstalled {
		fmt.Println("Intrerior already installed")
		icr.next.execute(ccr)
		return
	}
	fmt.Println("Installation interior by workers")
	ccr.interiorInstalled = true
	icr.next.execute(ccr)
}

func (icr *interiorCR) setNext(next assemblingLineCR) {
	icr.next = next
}

type wheelsCR struct {
	next assemblingLineCR
}

func (wcr *wheelsCR) execute(ccr *carCR) {
	if ccr.wheelsInstaled {
		fmt.Println("Wheels already installed")
	}
	fmt.Println("Installation wheels by workers")
}

func (wcr *wheelsCR) setNext(next assemblingLineCR) {
	wcr.next = next
}

/*func main() {

	car := &carCR{
		model: "new",
	}

	bodyShop := &bodyShopCR{}
	engineInstallation := &engineInstallationCR{}
	interior := &interiorCR{}
	wheelsCR := &wheelsCR{}

	bodyShop.setNext(engineInstallation)
	engineInstallation.setNext(interior)
	interior.setNext(wheelsCR)

	bodyShop.execute(car)

}*/
