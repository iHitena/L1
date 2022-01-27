package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {

	str := "snow dog sun 言葉"

	fmt.Print(str)
	var result string = " —"
	// создаем срез из строки . разделяя элементы символом - пробелом
	slice := strings.Split(str, " ")

	countWord := cap(slice) - 1
	// достаем элементы начиная с хвоста
	for i := countWord; i > -1; i-- {
		result += " " + slice[i]
	}

	fmt.Println(result)
}
