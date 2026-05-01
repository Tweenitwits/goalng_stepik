/*
Даны два числа. Найти их среднее арифметическое.

Формат входных данных
На вход дается два целых положительных числа a и b.

Формат выходных данных
Программа должна вывести среднее арифметическое чисел a и b
(ответ может быть целым числом или дробным)
*/

package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(err)
		return
	}
	arithmeticMean := (float64(a) + float64(b)) / 2
	fmt.Println(arithmeticMean)
}
