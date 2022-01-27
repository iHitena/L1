package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Реализовать бинарный поиск встроенными методами языка.

https://pkg.go.dev/sort#Search
*/

func main() {

	str := "Valera, what is you name ?"
	strSlice := strings.Split(str, " ")
	//бинарный поиск осуществляет в отсортированном массиве.
	//так что необходимо отсортировать, прежде чем выводить ответ

	//сортируем
	sort.Slice(strSlice, func(i, j int) bool {
		return strSlice[i] < strSlice[j]
	})

	//	производим пбинарный оиск
	resultStrSlise := sort.SearchStrings(strSlice, "name")

	fmt.Println(strSlice)
	fmt.Println(strSlice[resultStrSlise], resultStrSlise)

	intSlice := []int{3, 4, 5, 6, 1, 3, 5, 67, 4}

	//сортируем
	sort.Slice(intSlice, func(i, j int) bool {
		return intSlice[i] < intSlice[j]
	})

	//	производим бинарный поиск
	resultIntSlice := sort.SearchInts(intSlice, 4)

	fmt.Println(intSlice)
	fmt.Println(intSlice[resultIntSlice], resultIntSlice)

}
