/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.

Sample Input:
1234
Sample Output:
1
*/
package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	var a int = i / 10000
	var b int = i / 1000
	var c int = i / 100
	var d int = i / 10
	var e int = i / 1
	switch {
	case i == 10000:
		fmt.Println(1)
	case 0 < a:
		fmt.Println(a)
	case 0 < b:
		fmt.Println(b)
	case 0 < c:
		fmt.Println(c)
	case 0 < d:
		fmt.Println(d)
	case 0 < e:
		fmt.Println(e)
	}
}
