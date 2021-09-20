package main

/*
Функция «or» объединяет несколько каналов в один, если один (или все закрыты).
В качестве аргумента, функция получает:
	1) Срез каналов типа <-chan interface{}, который является исходным
	для последующего объединения его элементов в один канал;
	2) Канал типа chan interface{} – это инициализированный канал,
	который будет выдан в качестве результата работы функции.
Возвращаемое значение – канал типа <-chan interface{}.
Порядок работы функции: перебираем полученный срез типа <-chan interface{}
и для каждого его элемента в цикле от 0 до значения == длине этого элемента (канала)
и записываем полученный результат в канал для выдачи.
*/

func or(channels []<-chan interface{}, forReturned chan interface{}) <-chan interface{} {

	for _, oneChannel := range channels {

		for i := 0; i < len(oneChannel); {
			valueFromChannel := <-oneChannel
			forReturned <- valueFromChannel
		}
	}

	return forReturned
}

func main() {
}
