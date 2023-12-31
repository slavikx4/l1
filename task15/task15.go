package main

import (
	"math/rand"
)

// глобальная переменная
var justString string

// функция генерирующая строку заданой длины
func createHugeString(n int) string {
	// берём не только аски таблицу, чтобы показать реализацию проблемы
	alphabet := []rune("абвгдеёжзийклмнопрстуфхцчшщыьъэюяabcdefghijklmnopqrstuvwxyz0123456789")
	// длина алфавита, чтобы в следующем шаге не считать каждый раз заново
	lenAlphabet := len(alphabet)

	// результативный массив
	res := make([]rune, n)

	// инициализируем массив
	for i := 0; i < n; i++ {
		res[i] = alphabet[rand.Intn(lenAlphabet)]
	}

	// возвращаем строку
	return string(res)
}

// функция инициализирующая глобальную переменную
func someFunc() {
	// вызов функции, которая возвращат строку заданой длины
	v := createHugeString(1 << 10)
	// результативный слайс рун, который и будет в дальнейшем подстрокой(руны нужны для представления символов utf-8)
	res := make([]rune, 100)
	// проходимся по каждому символу строки
	for i, el := range v {
		// 100 - длина искомой подстроки
		if i < 100 {
			res[i] = el
		} else {
			break
		}
	}

	// кастуем слайс рун в строку
	justString = string(res)
}

// основной поток
func main() {

	// вызов функции, которая инициализирует глобальную переменную
	someFunc()
}
