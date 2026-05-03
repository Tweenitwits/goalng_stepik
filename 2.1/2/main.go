/*
Напишите функцию,
находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные

Вводится четыре числа.

Выходные данные

Необходимо вернуть из функции наименьшее из 4-х данных чисел
*/

package main

import "fmt"

func minimumFromFour() int {
	var minValue int
	first := true
	for i := 1; i <= 4; i++ {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if first {
			minValue = num
			first = false
		}
		if num < minValue {
			minValue = num
		}
	}
	return minValue
}

func main() {
	fmt.Println(minimumFromFour())
}
