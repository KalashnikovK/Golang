/*
Количество минимумов
Найдите количество минимальных элементов в последовательности.

Входные данные

Вводится натуральное число N, а затем N целых чисел последовательности.

Выходные данные

Выведите количество минимальных элементов последовательности.

Sample Input:
3
21
11
4

Sample Output:
1
*/

package main

import "fmt"

func main() {
	var dataLen int
	fmt.Scan(&dataLen)

	data := make([]int, dataLen)
	for idx, _ := range data {
		fmt.Scan(&data[idx])
	}

	min := data[0]
	amount := 0
	for _, value := range data {
		if value < min {
			min = value
			amount = 1
		} else if value == min {
			amount++
		}
	}

	fmt.Print(amount)
}
