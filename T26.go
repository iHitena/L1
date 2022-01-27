package main

import "fmt"

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
	aabcd — false
*/

func main() {
	a := "ddssdsdDD"

	b := "avbth"

	fmt.Println(rep(a))

	fmt.Println(rep(b))

}

func rep(str string) bool {
	m := make(map[int32]int)
	// создаем ключи в карте, если ключ посторится завершим интерацию с ложью,
	// если цикл пробежит весь массив без повторения функция завершится с правдой
	for _, numSimb := range str {
		if _, flag := m[numSimb]; !flag {
			m[numSimb] = 1
		} else {
			return false
		}
	}
	return true
}
