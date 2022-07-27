/*
Дано трехзначное число. Найдите сумму его цифр.

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.

Sample Input:
745

Sample Output:
16
*/

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	res := 0
	for num != 0 {
		res += num % 10
		num /= 10
	}

	fmt.Print(res)
}
