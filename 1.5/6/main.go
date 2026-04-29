/*
Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.

Циферблат часов стандартный 12-ти часовой:

Входные данные

На вход программе подается целое число d (0 < d < 360).

Выходные данные

Выведите на экран фразу:

It is ... hours ... minutes.

Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.
*/

package main

import "fmt"

func main() {
	const hourInDegrees = 30
	var degree int
	_, err := fmt.Scan(&degree)
	if err != nil {
		return
	}
	hours := degree / hourInDegrees
	minutes := degree % hourInDegrees * 2
	fmt.Printf("It is %d hours %d minutes\n", hours, minutes)
}
