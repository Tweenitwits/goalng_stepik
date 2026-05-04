/*
Дается строка. Нужно удалить все символы,
которые встречаются более одного раза и вывести получившуюся строку
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	totalStr := ""
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range str {
		if strings.Count(str, string(v)) > 1 {
			continue
		}
		totalStr += string(v)
	}
	fmt.Println(totalStr)
}
