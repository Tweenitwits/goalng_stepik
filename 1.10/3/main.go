/*
Напишите программу,
которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
Программа в первой строке получает на вход число n -
количество чисел в последовательности, во второй строке -- n чисел,
входящих в данную последовательность.
*/

package main

import "fmt"

func main() {
	var n, sum int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for i := 0; i < n; i++ {
		var num int
		_, err2 := fmt.Scan(&num)
		if err2 != nil {
			fmt.Println("Error: ", err2)
			return
		}
		if num > 9 && num < 100 && num%8 == 0 {
			sum += num
		}
	}
	fmt.Println(sum)
}
