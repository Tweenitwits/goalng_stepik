/*
Входные данные

Даны три натуральных числа a, b, c. Определите,
существует ли треугольник с такими сторонами.

Выходные данные

Если треугольник существует, выведите строку "Существует",
иначе выведите строку "Не существует". Строку выводите без кавычек.
*/

package main

import "fmt"

func main() {
	var a, b, c int
	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		fmt.Println(err)
		return
	}

	firstCon := a+b > c
	secondCon := b+c > a
	thirdCon := a+c > b

	if firstCon && secondCon && thirdCon {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}
