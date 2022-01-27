package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func main() {
	fmt.Println("Statr sleep")

	// испозуем для сравнения родной слип
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}()

	testSleep(2)
	fmt.Println("End sleep")
}

func testSleep(second int) {
	startTime := time.Now()

	for {
		// циклим программу, пока время старта и необходимое количество секунд не станут меньше настоящего времени,
		// после выходим из цикла
		endTime := startTime.Add(time.Duration(second) * time.Second)

		if time.Now().Sub(endTime) >= 0 {
			return
		}
	}
	fmt.Println(startTime)

}
