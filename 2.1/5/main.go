/*
Напишите функцию sumInt,
принимающую переменное количество аргументов типа int,
и возвращающую количество полученных функцией аргументов и их сумму.
Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:

a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1
*/

package main

import "fmt"

func sumInt(nums ...int) (int, int) {
	sum := 0
	count := 0
	for _, number := range nums {
		sum += number
		count++
	}
	return count, sum
}

func main() {
	fmt.Print(sumInt(1, 2, 3, 4, 5))
}
