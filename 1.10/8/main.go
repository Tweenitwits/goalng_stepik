/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.

Входные данные

Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.

Выходные данные

Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом числе.
*/

package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		log.Fatal("Ошибка при считывании числа: ", err)
	}
	firstStr := strconv.Itoa(a)
	secondStr := strconv.Itoa(b)
	for _, char := range firstStr {
		char -= '0'
		for _, chr := range secondStr {
			chr -= '0'
			if char == chr {
				fmt.Print(int(char))
				fmt.Print(" ")
			}
		}
	}
}
