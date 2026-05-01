/*
Двоичная запись
Дано натуральное число N. Выведите его представление в двоичном виде.

Входные данные

Задано единственное число N

Выходные данные

Необходимо вывести требуемое представление числа N.
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
	fmt.Printf("%b", num)
}
