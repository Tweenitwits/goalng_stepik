/*
На вход подаются a и b - катеты прямоугольного треугольника.
Нужно найти длину гипотенузы
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(err)
	}
	hypotenuse := math.Sqrt(a*a + b*b)
	fmt.Println(hypotenuse)
}
