/*
Количество нулей
По данным числам, определите количество чисел, которые равны нулю.

Входные данные
Вводится натуральное число N, а затем N чисел.

Выходные данные
Выведите количество чисел, которые равны нулю.

Sample Input:
5
1
8
100
0
12

Sample Output:
1
*/

package main

import "fmt"

func main() {
	var dataLen int
	fmt.Scan(&dataLen)

	data := make([]int, dataLen)
	amount := 0
	for _, value := range data {
		fmt.Scan(&value)

		if value == 0 {
			amount += 1
		}
	}

	fmt.Print(amount)
}
