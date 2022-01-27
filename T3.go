package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
с использованием конкурентных вычислений.

*/

func main() {
	// создаем массив и заполняем его
	array := [5]int{2, 4, 6, 8, 10}

	countArray := len(array)
	// создаем канал для квадратов - ббуфер канал (есть размер)
	squareCh := make(chan int, countArray)
	//создаем канал для передачи значения

	resultCh := make(chan int)

	// запусскаем горутину подсчета суммы - получает значения с помощью канала квадратов,
	//вернет значения с помощью канала результатов
	go summ(squareCh, resultCh)

	wg := sync.WaitGroup{}
	wg.Add(countArray)

	for i := range array {
		go func(i int) {
			// отправляет данные с квадратом в канал где данные будут обрабатываться горутиной подсчета суммы
			squareCh <- array[i] * array[i]
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(<-resultCh)
}

func summ(sqChan <-chan int, resultChan chan<- int) {
	result := 0
	// создает мутикс для блокировки данных результата (что бы данные были корректны)
	m := sync.Mutex{}

	wg := sync.WaitGroup{}

	wg.Add(cap(sqChan))
	//  цикл работает на основании размерности буффер. канала
	for i := 0; i < cap(sqChan); i++ {
		go func(i int) {
			// производим блокировку для корректного взаимодействия с результативной переменной
			m.Lock()
			result += <-sqChan
			m.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	resultChan <- result
}
