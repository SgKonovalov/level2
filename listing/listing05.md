5. Что выведет программа? Объяснить вывод программы.

package main

type customError struct {
     msg string
}

func (e *customError) Error() string {
    return e.msg
}

func test() *customError {
     {
         // do something
     }
     return nil
}

func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
_______________________________________________________________________________________________________________
ОТВЕТ:

ПРОГРАММА ВЫВЕДЕТ:
error

ПОЯСНЕНИЕ:
1) Функция test() возвращает *customError => после её работы тип переменной определён + как результат,
у нас сформирован адрес в памяти. А указанный nil в return – это то,
что будет храниться по этому адресу (неважно nil или существующий объект);
2) Создавая var err error мы задаём тип переменной err = error, со значением равным nil.
После отработки метода err = test(), err становится указателем на область памяти,
содержащий *customError (а это уже nil), а поэтому err не может стать nil.

РЕШЕНИЕ:
1) Создаём указатель на тип customError;
2) При помощи функции test(), записываем значение nil по адресу этой переменной:

package main

type customError struct {
    msg string
}

func (e *customError) Error() string {
    return e.msg
}

func test() *customError {
    {
        // do something
    }
    return nil
}

func main() {
    var err = (new(customError))
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}