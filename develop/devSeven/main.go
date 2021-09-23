package main

import (
	"fmt"
	"time"
)

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

func orV1(channels []<-chan interface{}, forReturned chan interface{}) <-chan interface{} {

	for _, oneChannel := range channels {

		for i := 0; i < len(oneChannel); {
			valueFromChannel := <-oneChannel
			forReturned <- valueFromChannel
		}
	}

	return forReturned
}

func orV2(channels ...<-chan interface{}) <-chan interface{} {

	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-orV2(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	var x []<-chan interface{}
	x = append(x, sig(2*time.Hour))
	x = append(x, sig(5*time.Minute))
	x = append(x, sig(1*time.Second))
	x = append(x, sig(1*time.Hour))
	x = append(x, sig(1*time.Minute))
	z := make(chan interface{})

	start1 := time.Now()
	<-orV1(x, z)

	start := time.Now()
	<-orV2(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))
	fmt.Printf("fone after %v", time.Since(start1))

}
