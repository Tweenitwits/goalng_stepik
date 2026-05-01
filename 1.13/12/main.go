/*
По данному числу N распечатайте все целые значения степени двойки,
не превосходящие N, в порядке возрастания.

Входные данные

Вводится натуральное число.

Выходные данные:
Выведите ответ на задачу.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var maxValue int
	num := 0
	_, err := fmt.Scan(&maxValue)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		tmp := math.Pow(2, float64(num))
		if tmp > float64(maxValue) {
			break
		} else {
			fmt.Print(tmp)
			fmt.Print(" ")
			num++
		}
	}
}
