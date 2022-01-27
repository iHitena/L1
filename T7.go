package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

func main() {

	array := [11]int{1, 2, 3, 4, 5, 6, 1, 8, 2, 10, 11}
	//инициализируем карту
	m := make(map[int]int)

	rvm := sync.RWMutex{}
	wg := sync.WaitGroup{}

	for _, value := range array {
		//добавляем потом в группу
		wg.Add(1)

		go func(m map[int]int, value int) {
			//блокируем запись во время изменения переменной
			rvm.Lock()
			//	создаем новый ключ со значением в мапе если его не было ранее или изменяет значение
			if _, flag := m[value]; flag {

				m[value]++
				rvm.Unlock()
			} else {

				m[value] = 1
				rvm.Unlock()
			}

			wg.Done()
		}(m, value)

	}
	wg.Wait()
	fmt.Println(m)
}
