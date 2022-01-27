package main

import (
	"fmt"
	"sort"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	array := []int{10, 2, 8, 7, 6, 5, 5, 4, 6, 7, 3}

	fmt.Println(array)
	// https://pkg.go.dev/sort#Slice
	sort.Slice(array, func(i, j int) bool {
		//задаем правила для сортировки
		return array[i] < array[j]
	})

	fmt.Println(array)
}
