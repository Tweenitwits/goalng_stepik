/*
Определите, является ли билет счастливым. Счастливым считается билет,
в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.

Формат входных данных

На вход подается номер билета - одно шестизначное число.

Формат выходных данных
Выведите "YES", если билет счастливый, в противном случае - "NO".
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
	firstDigit := (num / 100000) % 10
	secondDigit := (num / 10000) % 10
	thirdDigit := (num / 1000) % 10
	fourthDigit := (num / 100) % 10
	fifthDigit := (num / 10) % 10
	sixthDigit := num % 10
	firstSum := firstDigit + secondDigit + thirdDigit
	lastSum := fourthDigit + fifthDigit + sixthDigit
	if firstSum == lastSum {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
