/*
Напишите программу,
которая считывает целые числа с консоли по одному числу в строке.

Для каждого введённого числа проверить:

Если число меньше 10, то пропускаем это число;
если число больше 100, то прекращаем считывать числа;
в остальных случаях вывести это число обратно на консоль в отдельной строке.
*/

package main

import "fmt"

func main() {
loop:
	for {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		switch {
		case num < 10:
			continue
		case num > 100:
			break loop
		default:
			fmt.Println(num)
		}
	}
}
