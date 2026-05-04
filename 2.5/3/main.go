/*
Даются две строки X и S.
Нужно найти и вывести индекс первого вхождения подстроки S в строке X.
Если подстроки S нет в строке X - вывести -1
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var stringX, stringS string
	_, err := fmt.Scan(&stringX, &stringS)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.Index(stringX, stringS))
}
