/*
По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".
*/

package main

import "fmt"

func main() {
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	firstDigit := num / 100
	secondDigit := num % 100 / 10
	lastDigit := num % 10
	if firstDigit != secondDigit && secondDigit != lastDigit && firstDigit != lastDigit {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
