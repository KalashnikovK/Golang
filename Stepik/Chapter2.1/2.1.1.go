/*
Напишите функцию f(), которая будет принимать строку text и выводить ее (печатать на экране).

От вас требуется дописать только эту функцию, считайте что функция main() уже объявлена, считывать с консоли ничего не нужно!

Пакет "fmt" уже импортирован!

Sample Input:

Hello!
Sample Output:

Hello!
*/

package main

import "fmt"

func main() {
	text := "Hello!"

	f(text)
}

func f(text string) {
	fmt.Print(text)
}
