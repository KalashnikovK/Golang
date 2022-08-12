/*
Используем анонимные функции на практике.
Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint,
которую внутри функции main в дальнейшем можно будет вызвать по имени fn.

Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0.
Если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.

Пакет main объявлять не нужно!
\Вводить и выводить что-либо не нужно!
Уже импортированы - "fmt" и "strconv", другие пакеты импортировать нельзя.

Sample Input:
727178

Sample Output:
28
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num uint = 0000000

	// в степик нужна только эта функция, от сих
	fn := func(num uint) uint {
		str := strconv.Itoa(int(num))
		var inter string
		for _, value := range str {
			n, _ := strconv.Atoi(string(value))

			if n != 0 && n%2 == 0 {
				inter += strconv.Itoa(n)
			}
		}

		res, _ := strconv.Atoi(inter)

		if res == 0 {
			return 100
		}

		return uint(res)
	}
	// до сих

	fmt.Print(fn(num))
}
