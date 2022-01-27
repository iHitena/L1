package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает,
 вычитает две числовых переменных a,b, значение которых > 2^20.
*/

func main() {
	// используем библиотеку биг маф и создаем БОЛЬШИЕ переменные
	a := big.NewInt(10)
	b := new(big.Int)
	// присваевание
	b.SetString("1", 10)
	// сложение
	b.Add(b, b)

	for i := 0; i < 10; i++ {
		умножение
		a.Mul(a, a)

	}

	fmt.Println(a)
	// деление
	a.Quo(a, b)

	fmt.Println(b)

	fmt.Println(a)
	//вычитание
	b.Sub(b, big.NewInt(3))

	fmt.Println(b)

}
