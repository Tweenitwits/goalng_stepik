/*
Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных:
Выведите перевернутое число.
*/

package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
		return
	}
	reversed := 0
	for num != 0 {
		remainder := num % 10
		reversed = reversed*10 + remainder
		num /= 10
	}
	fmt.Println(reversed)
}
