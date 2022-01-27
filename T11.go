package main

import (
	"fmt"
	"strconv"
)

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {

	intersectionArray := make(map[string]int)

	array1 := []int{10, 10, 1, 1, 23, 4, 5, 6, 7, 8, 8, 9}
	array2 := []string{"10", "a", "b", "5", "7"}

	// добавляем ключи в карту со значением 1. повторный ключ не учитывается
	for _, value := range array1 {
		strValue := strconv.Itoa(value)
		if _, flag := intersectionArray[strValue]; !flag {
			intersectionArray[strValue] = 1
		}
	}
	// увеличиваем значение ключа, если он был найден во 2ром массиве
	for _, value := range array2 {
		if _, flag := intersectionArray[value]; flag {
			intersectionArray[value]++
		}
	}

	fmt.Println(intersectionArray)
	//	выводим ключи чь значения более 1 (т.е. онибыли замечены в 1 и 2 массиве)
	for key, count := range intersectionArray {
		if count > 1 {
			fmt.Print(key, " ")
		}
	}
}
