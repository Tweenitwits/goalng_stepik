/*
По данному числу n закончите фразу "На лугу пасется..."
одним из возможных продолжений: "n коров", "n корова",
"n коровы", правильно склоняя слово "корова".

Входные данные

Дано число n (0<n<100).

Выходные данные

Программа должна вывести введенное число n и одно из слов (на латинице):
korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov.
Между числом и словом должен стоять ровно один пробел.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hasAnySuffix(s string, suffix []string) bool {
	for _, suffix := range suffix {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

func main() {
	var n int
	var message string

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
		return
	}

	str := strconv.Itoa(n)
	suffixes := []string{"2", "3", "4"}
	wrongNums := []string{"12", "13", "14"}
	found := false

	for _, elem := range wrongNums {
		if elem == str {
			found = true
			break
		}
	}

	if strings.HasSuffix(str, "1") && str != "11" {
		message = "korova"
	} else if hasAnySuffix(str, suffixes) && !found {
		message = "korovy"
	} else {
		message = "korov"
	}

	fmt.Printf("%d %s", n, message)
}
