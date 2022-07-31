/*
Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1. Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
Напишите функцию, которая по указанному натуральному n возвращает φn.

Входные данные

Вводится одно число n.

Выходные данные

Необходимо вывести  значение φn.

Sample Input:
4

Sample Output:
3
*/

package main

import "fmt"

func main() {
	fmt.Print(fibonacci(4))
}

func fibonacci(n int) int {
	a := 0
	b := 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}

	return b
}
