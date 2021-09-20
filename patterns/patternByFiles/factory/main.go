package main

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

import (
	"factory/factory"
	"fmt"
)

func main() {
	SUV, _ := factory.GetCarF("SUV")
	Cabrio, _ := factory.GetCarF("Cabrio")
	Coupe, _ := factory.GetCarF("Coupe")
	fmt.Printf("%s price is %d\n", SUV.GetType(), SUV.GetPrice())
	fmt.Printf("%s price is %d\n", Cabrio.GetType(), Cabrio.GetPrice())
	fmt.Printf("%s price is %d\n", Coupe.GetType(), Coupe.GetPrice())
}