/*
Напишите функцию, которая умножает значения на которые ссылаются два указателя и выводит получившееся произведение в консоль. Ввод данных уже написан за вас.

Sample Input:
2 2

Sample Output:
4
*/

package main

import "fmt"

func main() {
	a := 2
	b := 4

	tests(&a, &b)
}

func tests(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Print(*x1, " ", *x2)
}
