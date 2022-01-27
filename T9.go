package main

import (
	"fmt"
	"time"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
 во второй — результат операции x*2,
 после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	// создаем каналы для передачи обычных чисел и квадратных

	chanNumbers := make(chan int)
	chanSquare := make(chan int)
	arrayNumbers := []int{}
	//	добавляем элементы в массив
	for i := 1; i < 100; i++ {
		arrayNumbers = append(arrayNumbers, i)
	}

	go func() {
		// читаем из массива и передаем элементы в канал, после закрываем,
		// что бы горутины которые используют канал завершили свою работу
		for i := range arrayNumbers {

			chanNumbers <- arrayNumbers[i]
		}
		close(chanNumbers)
	}()

	go func() {
		//чидаем данные из канала и возводим в квартад, далее передаем их в канал .
		// горутина завершит свою работу когда канал чтения закроется и там закончатся элементы
		for number := range chanNumbers {
			chanSquare <- number * number
		}
		close(chanSquare)
	}()

	go func() {
		// читаем из канала и выводим на экран пока канал не закрою и эл не закончатся
		for sNumber := range chanSquare {
			fmt.Println(sNumber)
		}
	}()

	time.Sleep(5 * time.Second)
}
