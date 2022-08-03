/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов,
он должен содержать только цифры и буквы латинского алфавита. На вход подается строка-пароль.
Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:
fdsghdfgjsdDD1

Sample Output:
Ok
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password string
	fmt.Scan(&password)

	if len(password) < 5 {
		fmt.Print("Wrong password")
		return
	}

	for _, value := range password {
		if !unicode.Is(unicode.Latin, value) && !unicode.IsDigit(value) {
			fmt.Print("Wrong password")
			return
		}
	}

	fmt.Print("Ok")
}
