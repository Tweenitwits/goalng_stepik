/*
Самое большое число, кратное 7.
Найдите самое большее число на отрезке от a до b, кратное 7.

Входные данные:
Вводится два целых числа a и b (a≤b).

Выходные данные:
Найдите самое большее число на отрезке от a до b
(отрезок включает в себя числа a и b), кратное 7,
или выведите "NO" - если таковых нет.
*/

package main

import "fmt"

func main() {
	var a, b, maxi int
	first := true
	flag := false

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := a; i <= b; i++ {
		if first {
			maxi = i
			first = false
			if maxi%7 == 0 {
				flag = true
			}
		}

		if i > maxi && i%7 == 0 {
			maxi = i
			flag = true
		}
	}
	if flag {
		fmt.Println(maxi)
	} else {
		fmt.Println("NO")
	}
}

//func main() {
//	var a, b int
//	var slice []int
//
//	_, err := fmt.Scan(&a, &b)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	for i := a; i <= b; i++ {
//		if i%7 == 0 {
//			slice = append(slice, i)
//		}
//	}
//
//	if len(slice) != 0 {
//		fmt.Println(slices.Max(slice))
//	} else {
//		fmt.Println("NO")
//	}
//}
