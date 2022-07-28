/*
Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0.
Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).

Входные данные

Сначала задано число N — количество элементов в массиве (1≤N≤100). Далее через пробел записаны N чисел — элементы массива.
Массив состоит из целых чисел.

Выходные данные

Необходимо вывести все элементы массива с чётными индексами.

Sample Input:
6
4 5 3 4 2 3

Sample Output:
4 3 2
*/

package main

import "fmt"

func main() {
	var dataLen int
	fmt.Scan(&dataLen)

	data := make([]int, dataLen)
	answer := make([]int, 0)
	for idx, value := range data {
		fmt.Scan(&value)

		if (idx % 2) == 0 {
			answer = append(answer, value)
		}
	}

	for _, value := range answer {
		fmt.Print(value, " ")
	}
}

/*
Конечно в терминал насыпает сразу ответ после ввода числа с четным индексом. Но на платформе степик считает что оответ правильный.

package main

import "fmt"

func main() {
	var dataLen int
	fmt.Scan(&dataLen)

	data := make([]int, dataLen)
	for idx, value := range data {
		fmt.Scan(&value)

		if (idx % 2) == 0 {
      fmt.Print(value, " ")
		}
	}
}
*/
