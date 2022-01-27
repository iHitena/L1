package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба 繰り返し 言葉— 葉言 し返り繰 абырвалг»).
Символы могут быть unicode.
*/

func main() {

	var oldStr string
	var result string

	var str string

	go func() {
		for {
			result = ""
			// изменения происходят только в том случае, когда str изменяется
			if oldStr != str {
				for _, simb := range str {
					// записываем строку задом наперед
					result = string(simb) + result
				}

				fmt.Println(result)
				oldStr = str
			}
		}
	}()

	str = "главрыба 繰り返し 言葉"
	fmt.Println(str)

	time.Sleep(1 * time.Second)

	str = "credit"
	fmt.Println(str)

	time.Sleep(1 * time.Second)

}
