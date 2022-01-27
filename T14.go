package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	var a interface{}

	//запускаем горутину, которая свособна выводит тип переменной, пока программа работает
	go func() {
		for {
			fmt.Printf("%T\n", a)
			time.Sleep(time.Second / 2)
		}
	}()

	a = 1
	time.Sleep(1 * time.Second)
	a = "2"
	time.Sleep(1 * time.Second)
	a = true
	time.Sleep(1 * time.Second)

	ch := make(chan int)

	a = ch
	time.Sleep(1 * time.Second)
}
