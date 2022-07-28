/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".

Входные данные

Дано число n (0<n<100).

Выходные данные

Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.

Sample Input:
10

Sample Output:
10 korov
*/

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	var word string
	inter := 0
	if num > 4 && num < 21 {
		word = "korov"
	} else {
		inter = num % 10
	}

	switch inter {
	case 1:
		word = "korova"
	case 2, 3, 4:
		word = "korovy"
	default:
		word = "korov"
	}
	fmt.Print(num, " ", word)
}
