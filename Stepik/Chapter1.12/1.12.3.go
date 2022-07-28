/*
На ввод подаются пять целых чисел, которые записываются в массив. Однако эта часть программы уже написана.

Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.

Sample Input:
2
3
56
45
21

Sample Output:
56
*/
package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	switch {
	case array[0] > array[1] && array[0] > array[2] && array[0] > array[3] && array[0] > array[4]:
		fmt.Print(array[0])
	case array[1] > array[0] && array[1] > array[2] && array[1] > array[3] && array[1] > array[4]:
		fmt.Print(array[1])
	case array[2] > array[0] && array[2] > array[1] && array[2] > array[3] && array[2] > array[4]:
		fmt.Print(array[2])
	case array[3] > array[0] && array[3] > array[1] && array[3] > array[2] && array[3] > array[4]:
		fmt.Print(array[3])
	case array[4] > array[0] && array[4] > array[1] && array[4] > array[2] && array[4] > array[3]:
		fmt.Print(array[4])
	}
}
