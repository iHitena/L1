package main

import (
	"fmt"
	"strconv"
)

/*
Реализовать паттерн «адаптер» на любом примере.



Необходимо достичь результата, при котором, на вход подаются два числа, на выходе получаем строку из этих двух чисел.
Мы имеем:
клиента, который хранит 2 числа
сервис, который на вход принимает строковый срез, на выходе выдает строку, содержащую элементы среза
*/

type client struct {
	a, b int
}

//реализуем метод клиента, на хводе числа и интерфейс
func (c client) clientFunc(clInte clientInterface) {
	// вызов функции интерфейса
	fmt.Println(clInte.clientIntFunc(c.a, c.b))
}

type service struct {
}

// функция сервика которая преобразуем массив строк в строку
func (s service) servuceFunc(strs []string) string {
	var result string
	for _, str := range strs {
		result += str
	}
	return result
}

// клиентский интерфейс
type clientInterface interface {
	clientIntFunc(a, b int) string
}
type adapter struct {
	*service
}

// адаптер конвентирует два числа в строку из массивов, используя метод сервиса
func (a adapter) clientIntFunc(x, y int) string {

	strs := []string{strconv.Itoa(x), strconv.Itoa(y)}
	return a.servuceFunc(strs)

}

func main() {
	//создаем клиента
	cl := client{1, 2}
	// добавляем адаптер с сылкой на сервис
	adapt := adapter{service: &service{}}
	//	вызываем метод в клиенте, который в последствии воспользуется методов сервиса, используя конвентированные данные клиента
	cl.clientFunc(&adapt)

}
