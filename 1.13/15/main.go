/*
Из натурального числа удалить заданную цифру.

Входные данные

Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные

Вывести число без заданных цифр.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num, n int
	var slice []string

	_, err := fmt.Scan(&num, &n)
	if err != nil {
		fmt.Println(err)
		return
	}

	str := strconv.Itoa(num)

	for _, val := range str {
		if int(val)-'0' == n {
			continue
		} else {
			slice = append(slice, string(val))
		}
	}
	tmp := ""

	for _, chr := range slice {
		tmp += string(chr)
	}

	numb, err1 := strconv.Atoi(tmp)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(numb)
}
