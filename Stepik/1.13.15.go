/*
Из натурального числа удалить заданную цифру.

Входные данные

Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные

Вывести число без заданных цифр.

Sample Input:
38012732
3

Sample Output:
801272
*/

package main

import "fmt"

func main() {
	var number, num int
	fmt.Scan(&number, &num)

	data := make([]int, 0)
	for ; number > 0; number /= 10 {
		if number%10 != num {
			data = append(data, number%10)
		}
	}

	for i := len(data) - 1; i >= 0; i-- {
		fmt.Print(data[i])
	}
}
