package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {

	sizeWorker := 0
	fmt.Scanln(&sizeWorker)

	//создаем канал для воркеров (будут читать из общего)
	chData := make(chan int)

	//инициализируем воркеры
	for i := 0; i < sizeWorker; i++ {
		go startWorker(i, chData)
	}
	//кидаем в канал данные
	for i := 0; i < 100000000; i++ {

		chData <- i //"number " + string(i)

	}
	//закрываем канал
	close(chData)

	time.Sleep(10 * time.Second)

}

func startWorker(workerNum int, in <-chan int) {
	//  завершаются т.к. канал закрывается и были считаны все данные
	for i := range in {
		fmt.Println(workerNum, " Data: ", i)
		//ПОзволяет передавать поток
		runtime.Gosched()
	}
}
