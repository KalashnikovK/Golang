/*
Внутри функции main (функцию объявлять не нужно),
вам необходимо в отдельных горутинах вызвать функцию work() 10 раз и дождаться результатов выполнения вызванных функций.

Функция work() ничего не принимает и не возвращает. Пакет "sync" уже импортирован.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			work1(i) // в степик нужно отдать work(), то есть изменить название и убрать i
			wg.Done()
		}()
	}

	wg.Wait()
}

func work1(n int) {
	fmt.Println(n)
}
