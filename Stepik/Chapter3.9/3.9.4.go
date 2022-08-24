/*
 Вам необходимо написать функцию calculator следующего вида:

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int
В качестве аргумента эта функция получает два канала только для чтения, возвращает канал только для чтения.

Через канал arguments функция получит ряд чисел, а через канал done - сигнал о необходимости завершить работу.
Когда сигнал о завершении работы будет получен, функция должна в выходной (возвращенный) канал отправить сумму полученных чисел.

Функция calculator должна быть неблокирующей, сразу возвращая управление.

Выходной канал должен быть закрыт после выполнения всех оговоренных условий, если вы этого не сделаете, то превысите предельное время работы.
*/

package main

import (
	"fmt"
)

// в степик нужна только эта функция

func calculator1(arguments <-chan int, done <-chan struct{}) <-chan int { // нужно только ее переименовать в calculator, без 1 на конце
	ch := make(chan int)

	go func() {
		defer close(ch)
		res := 0

		for {
			select {
			case x := <-arguments:
				res += x
			case <-done:
				ch <- res
				return
			}
		}
	}()

	return ch
}

// до сюда

func main() {
	arg := make(chan int)
	done := make(chan struct{})

	res := calculator1(arg, done)
	arg <- 2
	arg <- 4
	arg <- 5
	done <- struct{}{}
	fmt.Print(<-res)
}
