/*
Внутри функции main (функцию объявлять не нужно),
вам необходимо в отдельной горутине вызвать функцию work() и дождаться результатов ее выполнения.

Функция work() ничего не принимает и не возвращает.
*/

package main

import "fmt"

func work() {
	fmt.Print("ok")
}

func main() {
	ch := make(chan struct{})

	go func() {
		work()
		close(ch)
	}()

	<-ch
}
