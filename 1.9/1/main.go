/*
На ввод подается целое число. Если число положительное -
вывести сообщение "Число положительное", если число отрицательное -
"Число отрицательное". Если подается ноль - вывести сообщение "Ноль".
Выводить сообщение без кавычек.
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
	switch {
	case num > 0:
		fmt.Println("Число положительное")
	case num < 0:
		fmt.Println("Число отрицательное")
	default:
		fmt.Println("Ноль")
	}
}
