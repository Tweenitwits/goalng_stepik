/*
Количество нулей
По данным числам, определите количество чисел, которые равны нулю.

Входные данные:
Вводится натуральное число N, а затем N чисел.

Выходные данные:
Выведите количество чисел, которые равны нулю.
*/

package main

import (
	"fmt"
)

func main() {
	var n int
	count := 0
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < n; i++ {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println(err)
			return
		}
		if num == 0 {
			count++
		}
	}
	fmt.Println(count)
}
