/*
На ввод подаются пять целых чисел, которые записываются в массив.
Однако эта часть программы уже написана. Вам нужно написать фрагмент кода,
с помощью которого можно найти и вывести максимальное число в этом массиве.
*/

package main

import "fmt"

func main() {
	array := [5]int{}
	var a int

	for i := 0; i < 5; i++ {
		_, err := fmt.Scan(&a)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		array[i] = a
	}

	maxi := array[0]

	for _, elem := range array {
		if elem > maxi {
			maxi = elem
		}
	}
	fmt.Println(maxi)
}
