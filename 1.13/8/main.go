/*
Количество минимумов

Найдите количество минимальных элементов в последовательности.

Входные данные

Вводится натуральное число N, а затем N целых чисел последовательности.

Выходные данные

Выведите количество минимальных элементов последовательности.
*/

package main

import "fmt"

func main() {
	var n, mini int
	count := 0
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 1; i <= n; i++ {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println(err)
			return
		}

		if i == 1 {
			mini = num
			count = 1
			continue
		}

		if num < mini {
			mini = num
			count = 1
		} else if num == mini {
			count++
		}

	}
	fmt.Println(count)
}
