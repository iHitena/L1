package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в
структуре Action от родительской структуры Human (аналог наследования). - Композиция структур
*/

// Создаем хуман с использованием внедрения структур
type Human struct {
	nameMetod
	speed
}

type nameMetod string

type speed struct {
	minSpeed, maxSpeed int
}

// создаем пару методов структуры хуман
func (H Human) humanMetod() nameMetod {
	println("Human metod: " + H.nameMetod)
	return H.nameMetod
}

func (H Human) averageSpeed() int {
	return (H.maxSpeed + H.minSpeed) / 2
}

// создаем структуру Action
type Action struct {
}

func (a Action) actionMetod(name nameMetod) nameMetod {
	//встраиваем один метод в другой
	// принимаем name  в виде параметра и передаем его в структуру хуман далее вызываем метод из струкруты хуман
	return Human{nameMetod: name}.humanMetod() + " in action metod"
}

func (a Action) averageRunSpeed(s speed) int {
	return Human{speed: s}.averageSpeed() * 2
}

func main() {
	//создаем экземпляр структуры Action
	action := Action{}

	// создаем экземпляр структуры, находящейся внутри хуман
	var nMetod nameMetod = "Wow"
	//вызываем метод внутри Action,который в свою очередь реализован совместно с встраиваемым методом
	println(action.actionMetod(nMetod))

	hSpee := speed{1, 5}

	fmt.Println(action.averageRunSpeed(hSpee))

}
