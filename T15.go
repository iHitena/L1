package main

import (
	"fmt"
	"strconv"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

*/

func main() {

	someFunc()
	fmt.Println()
	fmt.Println(justString)
}

var justString string

// начальная проблема заключалась в том, что глобальная переменная заимствовала срез у локальной.
// при этом локальная переменная занимала много памяти и не очищалась из нее, т.к. срези использовал ее малу часть.
// было принято решения копировать часть среза
func someFunc() {
	v := createHugeString()
	fmt.Println(v)
	byteStr := make([]byte, 100)

	// копируем часть среза
	copy(byteStr, v[:100])

	justString = string(byteStr)
}

func createHugeString() string {
	var str string

	for i := 1000; i > 1; i-- {
		str += strconv.Itoa(i / 100)
	}
	return str
}
