/*
В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции, как объекты из реальной жизни.
Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу.

Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.

Если значение On == false, то оба метода вернут false.

Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет, то метод вернет false.
Метод RideBike работает также, но только зависит от свойства Power.

Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой структуры с именем testStruct в функции main, в дальнейшем программа проверит результат.

Закрывающая фигурная скобка в конце main() вам не видна, но она есть.

Пакет main объявлять не нужно!

Удачи!

Sample Input:

Sample Output:
*/

package main

func main() {
	testStruct := &Bot{}
	testStruct.Shoot() // это в Степик не нужно, просто чтоб не ругался что testStruct не используется
}

type Bot struct {
	On    bool
	Ammo  int
	Power int
}

func (s *Bot) Shoot() bool {

	if s.On == false {
		return false
	}

	if s.Ammo > 0 {
		s.Ammo = s.Ammo - 1
		return true
	} else {
		return false
	}
}

func (s *Bot) RideBike() bool {
	if s.On == false {
		return false
	}

	if s.Power > 0 {
		s.Power = s.Power - 1
		return true
	} else {
		return false
	}

}
