/*
Паттерн «Строитель», является порождающим,
т.е. отвечает за удобное и безопасное создание новых объектов или даже целых семейств объектов.
Так, «Строитель»:
1)  Отделяет конструирование сложного объекта от его представления,
так что в результате одного и того же процесса конструирования могут получаться разные представления,
при помощи «разделения» создания объекта на конкретные шаги;
2) Конструирование объекта осуществляется внешними по отношению к нему сущностями,
называемыми строителями. Т.е. каждый такой строитель выполняет конструирование объекта по-своему (это шаги из п. 1)
3) Позволяет избавиться от конструктора со множеством опциональных параметров.
Пример: каждый автомобиль может иметь разную комплектацию:
1) Руль – кожаный или полиуретановый;
2) Сиденье – кожаное или тканевое;
3) Аудио – есть/нет;
4) Кондиционер – есть/нет.
Состав опций не меняется, а их свойства - разные =>
необходимо для каждой комплектации реализовать свой набор свойств при производстве автомобиля.
Так:
1) Создаём сам Builder (CarBuilder) – интерфейс, который будет «производить» автомобиль, сохраняя набор свойств, но меняя их «наполнение»;
2) Функция getCarBuilder() – позволяет «выбрать» нужного CarBuilder по ключевому слову;
3) Структура standartCar, набор её методов и конструктор – это конкретная реализация комплектации с конкртным набором свойств (реализует CarBuilder).
Аналогично и luxeCar
4) director – управляет Builderom. Получая сведение о конкретном CarBuilder, который нужен выдаёт «задание» на производство конкретного автомобиля.

«Плюсы использования»:
1) Создаём объект по шагам;
2) Используем один и тот же код для разных объектов;

«Минусы использования»:
1) Привязка к конкретным структурам строителей;
2) Сложный код.
*/

package main

type CarBuilder interface {
	installStearingWheel()
	installSeats()
	installAudio()
	installAirConditioner()
	getCar() car
}

func getCarBuilder(builderType string) CarBuilder {
	if builderType == "standart" {
		return &standartCar{}
	}

	if builderType == "luxe" {
		return &luxeCar{}
	}
	return nil
}

type standartCar struct {
	stearingWheel  string
	seats          string
	audio          string
	airConditioner string
}

func newStandart() *standartCar {
	return &standartCar{}
}

func (sc *standartCar) installStearingWheel() {
	sc.seats = "Polyurethane stering wheel"
}

func (sc *standartCar) installSeats() {
	sc.seats = "Cloth seats"
}

func (sc *standartCar) installAudio() {
	sc.audio = "No audio"
}

func (sc *standartCar) installAirConditioner() {
	sc.airConditioner = "No air conditioner"
}

func (sc *standartCar) getCar() car {
	return car{
		stearingWheel:  sc.stearingWheel,
		seats:          sc.seats,
		audio:          sc.audio,
		airConditioner: sc.airConditioner,
	}
}

type luxeCar struct {
	stearingWheel  string
	seats          string
	audio          string
	airConditioner string
}

func newLuxeCarBuilder() *luxeCar {
	return &luxeCar{}
}

func (lc *luxeCar) installStearingWheel() {
	lc.seats = "Leather stering wheel"
}

func (lc *luxeCar) installSeats() {
	lc.seats = "Leather seats"
}

func (lc *luxeCar) installAudio() {
	lc.audio = "Luxe audio"
}

func (lc *luxeCar) installAirConditioner() {
	lc.airConditioner = "Climate-control"
}

func (lc *luxeCar) getCar() car {
	return car{
		stearingWheel:  lc.stearingWheel,
		seats:          lc.seats,
		audio:          lc.audio,
		airConditioner: lc.airConditioner,
	}
}

type car struct {
	stearingWheel  string
	seats          string
	audio          string
	airConditioner string
}

type director struct {
	carBuilder CarBuilder
}

func newDirector(cb CarBuilder) *director {
	return &director{
		carBuilder: cb,
	}
}

func (d *director) setCarBuilder(cb CarBuilder) {
	d.carBuilder = cb
}

func (d *director) produceCar() car {
	d.carBuilder.installStearingWheel()
	d.carBuilder.installSeats()
	d.carBuilder.installAudio()
	d.carBuilder.installAudio()
	d.carBuilder.installAirConditioner()
	return d.carBuilder.getCar()
}

/*func main() {

	standartCarProduce := getCarBuilder("standart")
	directorStandart := newDirector(standartCarProduce)
	standartCar := directorStandart.produceCar()

	luxeCarProduce := getCarBuilder("luxe")
	directorLuxe := newDirector(luxeCarProduce)
	luxeCar := directorLuxe.produceCar()

	fmt.Println(standartCar)
	fmt.Println(luxeCar)

}*/
