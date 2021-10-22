/*
Петя торопился в школу и неправильно написал программу, которая сначала находит квадраты двух чисел, а затем их суммирует. Исправьте его программу.

Sample Input:
2
2
Sample Output:
8
*/
package main

import "fmt"

func main() {

	var a int
	var b int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Scan(&b) // считаем переменную 'b' с консоли

	a = a * a
	b = b * b
	var c int = a + b
	fmt.Println(c)
}
