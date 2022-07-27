/*
Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.

Sample Input:
843

Sample Output:
348
*/

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	for num != 0 {
		res := num % 10
		num /= 10
		fmt.Print(res)
	}
}
