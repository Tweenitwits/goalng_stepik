/*
Заданы три числа - a, b, c (a<b<c)- длины сторон треугольника.
Нужно проверить, является ли треугольник прямоугольным.
Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"
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

	hypotenuseSquare := c * c
	cathetersSum := a*a + b*b

	if hypotenuseSquare == cathetersSum {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
