// Петя торопился в школу и неправильно написал программу,
// которая сначала находит квадраты двух чисел, а затем их суммирует.
// Исправьте его программу.

/*
package main
import "fmt"
func main(){

  var a int
  fmt.Scan(&a) // считаем переменную 'a' с консоли
  fmt.Scan(&b) // считаем переменную 'b' с консоли

  a = a * a
  b = b * 2
  c = a + b
  fmt.Println(c)
}
*/

package main

import "fmt"

func main() {

	var a, b int
	_, err1 := fmt.Scan(&a)
	if err1 != nil {
		fmt.Println("Error: ", err1)
		return
	}

	_, err2 := fmt.Scan(&b)
	if err2 != nil {
		fmt.Println("Error: ", err2)
		return
	}

	a *= a
	b *= b
	c := a + b
	fmt.Println(c)
}
