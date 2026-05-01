/*
Дано трехзначное число. Найдите сумму его цифр.

Формат входных данных:
На вход дается трехзначное число.

Формат выходных данных:
Выведите одно целое число - сумму цифр введенного числа.
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
	firstDigit := num / 100
	secondDigit := num / 10 % 10
	thirdDigit := num % 10
	sum := firstDigit + secondDigit + thirdDigit
	fmt.Println(sum)
}
