package main

import (
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {

	quit := make(chan bool)
	go func() {
		fmt.Println("start gor bool ch")
		i := 0

		//горутину можно завершить проверяе каждый раз булево значение. она завершится, когда будево значение станет правдой
		for {
			select {
			case <-quit:
				fmt.Println("end gor bool ch")
				return
			default:
				i++
				fmt.Println(i)
				time.Sleep(1 * time.Second)
			}
		}

	}()

	time.Sleep(3 * time.Second)

	quit <- true

	chTest := make(chan int)

	go Test(chTest)
	chTest <- 1
	chTest <- 2
	chTest <- 3
	// горутина завершится  после того как канал будет закрыт и опустошон
	close(chTest)
	fmt.Println("chTest Closed")

	chTest <- 4
}

func Test(ch chan int) {
	fmt.Println("Start test gor (close ch)")
	for x := range ch {

		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("end test got")
}
