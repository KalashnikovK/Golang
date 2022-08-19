/*
На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).

Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.

Sample Input:
13.03.2018 14:00:15,12.03.2018 14:00:15

Sample Output:
24h0m0s
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	str, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	strs := strings.Split(string(str), ",")

	firstDate, _ := time.Parse("02.01.2006 15:04:05", strs[0])
	secondDate, _ := time.Parse("02.01.2006 15:04:05", strs[1])

	if firstDate.After(secondDate) {
		fmt.Print(firstDate.Sub(secondDate))
	} else {
		fmt.Print(secondDate.Sub(firstDate))
	}
}
