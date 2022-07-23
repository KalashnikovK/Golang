/*
Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4), а затем NN целых чисел-элементов среза.
На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

Sample Input:
5
41 -231 24 49 6

Sample Output:
49
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	var dataLen, num int

	fmt.Scan(&dataLen)

	if dataLen < 4 {
		fmt.Printf("error datalen must be >= 4\n")
		os.Exit(1)
	}

	data := make([]int, dataLen)

	for i := 0; i < dataLen; i++ {
		fmt.Scan(&num)
		data[i] = num
	}

	fmt.Print(data[3])
}
