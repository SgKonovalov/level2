4. Что выведет программа? Объяснить вывод программы.

package main

import "sync"

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)

	ch := make(chan int)

	go func() {
		defer close(ch)
		//defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {

		println(n)
	}
	wg.Wait()

}
_______________________________________________________________________________________________________________
ОТВЕТ:
ПРОГРАММА ВЫВЕДЕТ:
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!

ПРОБЛЕМЫ:
1) Анонимная функция запущена в горутине, а main – не ждёт её завершения. Необходимо добавить WaitGroup;
2) Функция, записывающая данные в канал, не закрывает его => цикл, считывая данные из канала, ожидает поступления новых,
чем блокирует канал, т.к. горутина, «пишущая» в канал, больше ничего в него отправлять не собирается.

РЕШЕНИЕ:

package main

import "sync"

func main() {
    wg := sync.WaitGroup{}

    wg.Add(1)

    ch := make(chan int)

    go func() {
        defer close(ch)
        defer wg.Done()
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()

    for n := range ch {

        println(n)
    }
    wg.Wait()

}