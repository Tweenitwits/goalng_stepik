/*
На вход дается строка, из нее нужно сделать другую строку,
оставив только нечетные символы (считая с нуля)
*/

package main

import "fmt"

func main() {
	var str string
	totalStr := ""
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println(err)
		return
	}
	for idx, chr := range str {
		if idx%2 == 0 {
			continue
		}
		totalStr += string(chr)
	}
	fmt.Println(totalStr)
}
