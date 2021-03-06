package main

import "fmt"

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.


Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

*/

func main() {

	array := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//	создаем карту, где ключ будет является группой, а элементы будут состоят из числового массива
	group := make(map[int]*[]float32)

	for _, x := range array {
		//получаем  диапазон группы (группа+10)
		var numberGroup int = int(x) / 10 * 10
		// проверяем наличи группы. в противном случае добавляем ключ новой
		if _, flagGroup := group[numberGroup]; flagGroup {

			// внутри группа состоит из среза. если срез уже есть,
			// то добавляем новый элемент.в противном случае создаем новый срез
			oldSlice := *group[numberGroup]
			newSlice := append(oldSlice, float32(numberGroup))

			group[numberGroup] = &newSlice

		} else {
			slice := []float32{float32(numberGroup)}
			group[numberGroup] = &slice
		}
	}
	// вывод ключа - группа и значения группы - срез состоящих из всех значений массива,пренадлежах этой группе
	for key, value := range group {
		fmt.Println(key, " ", *value)
	}

}
