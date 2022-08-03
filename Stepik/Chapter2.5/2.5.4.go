/*
 На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

Sample Input:
ihgewlqlkot

Sample Output:
hello
*/

package main

import "fmt"

func main() {
	var text string
	fmt.Scan(&text)

	for i := 1; i < len(text); i += 2 {
		fmt.Printf("%c", text[i])
		//fmt.Print(text[i : i+1])
	}
}
