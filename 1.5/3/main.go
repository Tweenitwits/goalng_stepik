/*
По данному целому числу, найдите его квадрат.

Формат входных данных
На вход дается одно целое число.

Формат выходных данных
Программа должна вывести квадрат данного числа.
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
	num *= num
	fmt.Println(num)
}
