/*
Двоичная запись
Дано натуральное число N. Выведите его представление в двоичном виде.

Входные данные

Задано единственное число N

Выходные данные

Необходимо вывести требуемое представление числа N.

Sample Input:
12

Sample Output:
1100
*/

package main

import (
	"fmt"
)

/*
func main() {
	var num int
	fmt.Scan(&num)

	fmt.Printf("%b", num)
}

 Ну или так..........
*/
func main() {
	var num int
	fmt.Scan(&num)

	data := make([]int, 0)
	for num > 0 {
		data = append(data, num%2)
		num /= 2
	}

	for i := len(data) - 1; i >= 0; i-- {
		fmt.Print(data[i])
	}
}
