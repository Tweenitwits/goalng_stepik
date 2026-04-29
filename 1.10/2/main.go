/*
Требуется написать программу,
при выполнении которой с клавиатуры считываются два натуральных числа A и B
(каждое не более 100, A < B). Вывести сумму всех чисел от A до B включительно.
*/

package main

import "fmt"

func main() {
	var A, B, sum int
	_, err := fmt.Scan(&A, &B)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for i := A; i <= B; i++ {
		sum += i
	}
	fmt.Println(sum)
}
