package main

/*
Паттерн «Фасад», является структурным, т.е. отвечает за построение удобных в поддержке иерархий классов.
Т.е. «Фасад» - это простой интерфейс для работы со сложной подсистемой, содержащей множество классов,
а следовательно он определяет интерфейс более высокого уровня, который упрощает использование остальной подсистемы.
Пример: есть сложная подсистема автомобиля:
1) Руль;
2) Сиденье;
3) Аудио;
4) Кондиционер.
Все они являются независимыми друг от друга, но ими можно управлять в единой системе – «Автомобиле».
Так, мы можем повернуть руль, включить музыку и т.д..
Следовательно, тип «Автомобиль» - это интерфейса высокого уровня,
содержащий указатель на составные части всей системы (руль, сиденье и т.д.) => у него есть метод,
позволяющий использовать функции составляющих его систему структур – выключить кондиционер,
поднять спинку сиденья и т.д..
Такой подход позволяет использовать объект типа «Автомобиль» для доступа ко всем функциям его подсистемы.

«Плюсы использования»:
1) Упрощение доступа ко всем элементам сложной подсистемы;
2) Упрощение дальнейшей поддержки структур (при изменении логики, достаточно изменить работу членов иерархии).

«Минусы использования»:
1) Сложное в составлении (можно «забыть» добавить зависимость);
2) Не весь функции иерархии может быть доступен.
*/

import (
	"facade/facade"
	"fmt"
)

func main() {
	m := facade.NewCar()
	fmt.Println(m.Actions())
}
