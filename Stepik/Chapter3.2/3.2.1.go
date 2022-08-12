/*
Напишите функцию с именем convert, которая конвертирует входное число типа int64 в uint16.

Считывать и выводить ничего не нужно!

Считайте что функция main и пакет main уже объявлены!

Sample Input:

Sample Output:
*/

package main

import "fmt"

func main() {
	var num int64
	num = 1234
	fmt.Printf("Число %d с типом %T", convert(num), convert(num))
}

// в степик нужно передать только эту функцию от сих
func convert(num int64) uint16 {
	return uint16(num)
}

// до сих
