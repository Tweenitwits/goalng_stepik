/*
На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и
вывести получившееся число.
Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81.
Дальше 1. Единица в квадрате - 1. В итоге получаем 811181
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num, result int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
	}

	str := strconv.Itoa(num)
	tmp := strings.Split(str, "")
	resultStr := ""

	for _, v := range tmp {
		v, _ := strconv.Atoi(v)
		v *= v
		resultStr += strconv.Itoa(v)
	}

	result, _ = strconv.Atoi(resultStr)
	fmt.Println(result)
}
