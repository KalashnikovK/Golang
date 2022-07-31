/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:

a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1

Sample Input:

Sample Output:
*/

package main

import "fmt"

func main() {
	a, b := sumInt(1, 0, 3, 8, 12, 0, 7)
	fmt.Println(a, b)
}

func sumInt(num ...int) (a int, b int) {
	sum := 0
	for _, value := range num {
		sum += value
	}
	return len(num), sum
}
