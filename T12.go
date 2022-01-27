package main

import (
	"fmt"
)

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

type row struct {
	name  string
	array []string
}

func main() {

	array := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(array)
	//	создаем slice из  сруктур, у которых есть свое собственное множество , состоящее из текстового массива и названия строки
	resultArray := []row{}

	for _, str := range array {
		//заполняем срез
		newRow := row{name: str}
		resultArray = append(resultArray, newRow)

	}

	fmt.Println(resultArray)

}
