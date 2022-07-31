/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные

Вводится четыре числа.

Выходные данные

Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:
4 5 6 7

Sample Output:
4
*/

package main

import "fmt"

func main() {

}

func minimumFromFour() int {
	var min int
	fmt.Scan(&min)

	inter := 0
	for i := 0; i < 3; i++ {
		fmt.Scan(&inter)
		if inter < min {
			min = inter
		}
	}

	return min
}
