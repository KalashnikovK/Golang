/*
По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:
237
Sample Output 1:
YES
Sample Input 2:
117
Sample Output 2:
NO
*/
package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	var a = i / 100
	var b = i / 10 % 10
	var c = i % 10
	switch {
	case a != b && b != c && a != c:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
