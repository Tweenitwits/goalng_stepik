/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

Формат входных данных:
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных:
Выведите одно целое число - первую цифру заданного числа.
*/

package main

import "fmt"

func main() {
	var num, firstDigit int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	switch {
	case num < 10:
		firstDigit = num % 10
	case 10 < num && num < 100:
		firstDigit = num / 10
	case 100 < num && num < 1000:
		firstDigit = num / 100
	case 1000 < num && num < 10000:
		firstDigit = num / 1000
	case num == 10000:
		firstDigit = num / 10000
	}
	fmt.Println(firstDigit)
}
