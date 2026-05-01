/*
Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0.
Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).

Входные данные:

Сначала задано число N — количество элементов в массиве (1≤N≤100).
Далее через пробел записаны N чисел — элементы массива.
Массив состоит из целых чисел.

Выходные данные

Необходимо вывести все элементы массива с чётными индексами.
Элементы выводить через пробел.
*/

package main

import "fmt"

func main() {
	var slice []int
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for i := 1; i <= n; i++ {
		var elem int
		_, err := fmt.Scan(&elem)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		slice = append(slice, elem)
	}

	for idx, elem := range slice {
		if idx%2 == 0 {
			fmt.Print(elem)
			fmt.Print(" ")
		}
	}
}
