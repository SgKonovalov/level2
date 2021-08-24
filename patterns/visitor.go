/*
Паттерн «Посетитель», является поведенческим, т.е. решает задачи эффективного и безопасного взаимодействия
между объектами программы. Так, «Посетитель»:
1)  Позволяет добавлять в программу новые операции, не изменяя (почти) классы объектов,
над которыми эти операции могут выполняться;
2) Использует механизм двойной диспетчеризации т.е. объекты, которые мы передаём в параметры «Посетителю»,
сами ищут нужный метод.
Так:
1) Создаём общий интерфейс, позволяющий в котором определяем метод accept.
Этот метод и будет «встраивать» новый функционал;
2) Реализуем этот интерфейс;
3) Создаём интерфейс «Посетителя»: определяем методы, которые будут добавлять новый функционал;

«Плюсы использования»:
1) Возможность дополнить функционал, не меняя основного кода;
2) Инкапсуляция этого доп. функционала;
«Минусы использования»:
1) Невозможно использовать, если иерархия часто меняется.
*/

package main

import "fmt"

type carV interface {
	getType() string
	accept(visitor)
}

type standartCarV struct {
	price int
}

func (scv *standartCarV) accept(v visitor) {
	v.visitForStandartCarVPrice(scv)
}

func (scv *standartCarV) getType() string {
	return "Standart car Price"
}

type luxeCarV struct {
	price int
}

func (lcv *luxeCarV) accept(v visitor) {
	v.visitForLuxeCarVPrice(lcv)
}

func (c *luxeCarV) getType() string {
	return "Luxe car Price"
}

type visitor interface {
	visitForStandartCarVPrice(*standartCarV)
	visitForLuxeCarVPrice(*luxeCarV)
}

type fullPrice struct {
	fullPrice int
}

func (fp *fullPrice) visitForStandartCarVPrice(scv *standartCarV) {
	fmt.Printf("Standart Car is %d\n", scv.price+500)
}

func (fp *fullPrice) visitForLuxeCarVPrice(lcv *luxeCarV) {
	fmt.Printf("Luxe Car is %d\n", lcv.price+1500)
}

/*func main() {
	lcv := luxeCarV{price: 2000}
	scv := standartCarV{price: 300}

	fullPrice := &fullPrice{}

	lcv.accept(fullPrice)
	scv.accept(fullPrice)

}*/
