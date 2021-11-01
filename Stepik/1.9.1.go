/*
На ввод подается целое число. Если число положительное - вывести сообщение "Число положительное",
если число отрицательное - "Число отрицательное". Если подается ноль - вывести сообщение "Ноль". Выводить сообщение без кавычек.

Sample Input:
5
Sample Output:
Число положительное
*/
package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	switch {
	case i < 0:
		fmt.Print("Число отрицательное")
	case i == 0:
		fmt.Print("Ноль")
	case i > 0:
		fmt.Print("Число положительное")
	}

}
