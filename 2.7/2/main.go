/*
Дана строка, содержащая только английские буквы (большие и маленькие).
Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и
после последней символ ‘*’ добавлять не нужно).

Входные данные

Вводится строка ненулевой длины. Известно также,
что длина строки не превышает 1000 знаков.

Выходные данные

Вывести строку, которая получится после добавления символов '*'.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tmp := ""
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := strings.TrimSpace(scanner.Text())

	if text == "" {
		return
	}

	for idx, v := range text {
		if idx == len(text) {
			break
		}
		tmp += string(v) + "*"
	}
	tmp = strings.TrimRight(tmp, "*")
	fmt.Println(tmp)
}
