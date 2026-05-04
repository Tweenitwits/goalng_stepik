/*
На вход подается строка. Нужно определить, является ли она правильной или нет.
Правильная строка начинается с заглавной буквы и заканчивается точкой.
Если строка правильная - вывести Right иначе - вывести Wrong

Маленькая подсказка: fmt.Scan() считывает строку до первого пробела,
вы можете считать полностью строку за раз с помощью bufio:

text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return
	}

	var upper, dot bool
	br := []rune(text)

	if text == "" {
		fmt.Println("Wrong")
	} else {
		text = strings.TrimRight(text, "\n")
		upper = unicode.IsUpper(br[0])
		dot = strings.HasSuffix(text, ".")
		if upper && dot {
			fmt.Println("Right")
		} else {
			fmt.Println("Wrong")
		}
	}
}
