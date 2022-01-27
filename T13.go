package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	a := 1000
	b := 10

	fmt.Println(a, " ", b)
	//меняем местами
	a, b = b, a
	fmt.Println(a, " ", b)

}
