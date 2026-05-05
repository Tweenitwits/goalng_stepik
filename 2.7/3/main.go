/*
Дана строка, содержащая только арабские цифры.
Найти и вывести наибольшую цифру.

Входные данные

Вводится строка ненулевой длины. Известно также,
что длина строки не превышает 1000 знаков и
строка содержит только арабские цифры.

Выходные данные

Выведите максимальную цифру, которая встречается во введенной строке.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := strings.TrimSpace(scanner.Text())

	if text == "" {
		return
	}

	var maxValue int
	first := true

	for _, v := range text {
		if first {
			maxValue = int(v) - '0'
			first = false
		}

		if int(v)-'0' > maxValue {
			maxValue = int(v) - '0'
		}
	}
	fmt.Println(maxValue)
}
