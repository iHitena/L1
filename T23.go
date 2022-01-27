package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func main() {

	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	i := 4
	//	присваем срезу новый срез без И элемента
	//slice = append(slice[:i], slice[i+1:])

	fmt.Println(slice, i)
}
