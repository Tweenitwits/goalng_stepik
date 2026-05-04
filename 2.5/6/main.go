/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
Длина пароля должна быть не менее 5 символов,
он может содержать только арабские цифры и буквы латинского алфавита.
На вход подается строка-пароль. Если пароль соответствует требованиям -
вывести "Ok", иначе вывести "Wrong password"
*/

package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	var pass string
	_, err := fmt.Scan(&pass)
	if err != nil {
		fmt.Println(err)
		return
	}

	symbolCheck := true

	lengthCheck := utf8.RuneCountInString(pass) >= 5

	for _, v := range pass {
		if !unicode.IsDigit(v) && !unicode.Is(unicode.Latin, v) {
			symbolCheck = false
			break
		}
	}

	if symbolCheck && lengthCheck {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
