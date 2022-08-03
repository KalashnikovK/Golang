/*
На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").

Sample Input:
топот

Sample Output:
Палиндром
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\r\n")

	runes := []rune(text)
	t := 1
	for _, value := range runes {
		if value == runes[len(runes)-t] {
		} else {
			fmt.Print("Нет")
			return
		}
		t++
	}
	fmt.Print("Палиндром")
}
