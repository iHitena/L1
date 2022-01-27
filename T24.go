package main

import (
	"fmt"

	"./TP24"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными
параметрами x,y и конструктором
*/

func main() {
	// создаем тип из пакета (в пакете 3 переменные , 2из них приватные)
	point := TP24.NewPoint(2, 4)
	//вызываем мето, использующий приватные переменные в пакете
	distace := point.DistanceXY()

	fmt.Println(distace)

	fmt.Println(point.Distance)

	fmt.Println(*point)

}
