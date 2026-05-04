/*
На вход подается строка. Нужно определить, является ли она палиндромом.
Если строка является палиндромом - вывести Палиндром иначе - вывести Нет.
Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст,
одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := strings.TrimSpace(scanner.Text())

	if text == "" {
		fmt.Println("Нет")
		return
	}

	if isPalindrome(text) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
