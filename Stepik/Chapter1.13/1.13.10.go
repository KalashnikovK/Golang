/*
Самое большое число, кратное 7
Найдите самое большее число на отрезке от a до b, кратное 7 .

Входные данные
Вводится два целых числа a и b (a≤b).

Выходные данные
Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b), кратное 7 , или выведите "NO" - если таковых нет.

Sample Input:
100
500

Sample Output:
497
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Printf("error %d must be <= %d", a, b)
		os.Exit(1)
	}

	for ; a <= b; b-- {
		if b%7 == 0 {
			fmt.Print(b)
			return
		}
	}
	fmt.Print("NO")
}
