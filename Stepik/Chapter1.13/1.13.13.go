/*
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.

Входные данные

Вводится натуральное число.

Выходные данные

Выведите ответ на задачу.

Sample Input:
8

Sample Output:
6
*/

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	data := make([]int, 2)

	data[0] = 0
	data[1] = 1
	inter := 0
	for n := 2; inter < num; n++ {
		data = append(data, data[n-2]+data[n-1])
		inter = data[n]
		if inter == num {
			fmt.Print(len(data) - 1)
			return
		}
	}

	fmt.Print(-1)
}

/*
package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	counter := 1
	a := 0
	b := 1
	for b < num {
		counter++
		a, b = b, a+b
	}

	if num == b {
		fmt.Println(counter)
	} else {
		fmt.Print(-1)
	}
}
*/
