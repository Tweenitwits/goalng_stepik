/*
Напишите программу, принимающаю на вход число
N(N≥4), а затем N целых чисел-элементов среза.
На вывод нужно подать 4-й (3 по индексу) элемент данного среза.
*/

package main

import "fmt"

func main() {
	var slice []int
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for i := 0; i < n; i++ {
		var el int
		_, err = fmt.Scan(&el)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		slice = append(slice, el)
	}
	fmt.Println(slice[3])
}
