/*
Паттерн «Фабричный метод», является порождающим паттерном, т.е. отвечает за удобное и безопасное создание новых объектов или даже целых семейств объектов.
Он позволяет:
1) Определяет интерфейс для создания объекта, но оставляет подклассам решение о том, какой класс инстанцировать. Фабричный метод позволяет классу делегировать инстанцирование подклассам;
2) Отделяет код производства продуктов от остального кода, который эти продукты использует;
3) Создание объектов внутри класса с помощью фабричного метода всегда оказывается более гибким решением, чем непосредственное создание;

Пример: создание автомобиля по типу кузова.
1) Определяем интерфейс, инкапсулирующий методы по установке и выдаче значений для типа кузова;
2) Определяем «родительскую» структуру автомобиля (имеет имя и цену) + реализовываем интерфейс из п. 1;
3) Создаём сами типы автомобилей – структура с композицией из п. 2;
4) Функция getCarF – производит выдачу нужного типа, основываясь на его названии.

«Плюсы использования»:
1) Добавление новых объектов становится проще;
2) Не нужно привязываться к конкретным классам;
3)) Код для инициализации новых объектов в одном месте.

«Минусы использования»:
1) Код усложняется: может быть много разных иерархий.
*/

package main

import "fmt"

type bodyTypeF interface {
	setType(name string)
	setPrice(price int)
	getType() string
	getPrice() int
}

type carF struct {
	name  string
	price int
}

func (cf *carF) setType(name string) {
	cf.name = name
}

func (cf *carF) getType() string {
	return cf.name
}

func (cf *carF) setPrice(price int) {
	cf.price = price
}

func (cf *carF) getPrice() int {
	return cf.price
}

type SUV struct {
	carF
}

func newSUV() bodyTypeF {
	return &SUV{
		carF: carF{
			name:  "SUV",
			price: 100000,
		},
	}
}

type cabrio struct {
	carF
}

func newCabrio() bodyTypeF {
	return &SUV{
		carF: carF{
			name:  "Cabrio",
			price: 150000,
		},
	}
}

type coupe struct {
	carF
}

func newCoupe() bodyTypeF {
	return &SUV{
		carF: carF{
			name:  "Coupe",
			price: 200000,
		},
	}
}

func getCarF(bodyType string) (bodyTypeF, error) {
	if bodyType == "SUV" {
		return newSUV(), nil
	}
	if bodyType == "Cabrio" {
		return newCabrio(), nil
	}
	if bodyType == "Coupe" {
		return newCoupe(), nil
	}
	return nil, fmt.Errorf("Wrong body type passed")
}

/*func main() {
	SUV, _ := getCarF("SUV")
	Cabrio, _ := getCarF("Cabrio")
	Coupe, _ := getCarF("Coupe")
	fmt.Printf("%s price is %d\n", SUV.getType(), SUV.getPrice())
	fmt.Printf("%s price is %d\n", Cabrio.getType(), Cabrio.getPrice())
	fmt.Printf("%s price is %d\n", Coupe.getType(), Coupe.getPrice())
}*/
