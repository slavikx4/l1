package main

import (
	"fmt"
	"unicode"
)

// проверяет состоит ли строка из уникальных символов
func uniqueChars(s string) bool {

	// записываем все символы в мапу
	// в ключ щаписываем символ
	// а в ззначение ставим "заглушку"
	// записываем в руне, чтобы читать юникод
	set := make(map[rune]struct{})
	for _, r := range s {

		// приводим к нижнему регистру
		r = unicode.ToLower(r)

		// проверка на наличие сивола в мапе,
		// если есть, то прекращаем функцию и возвращаем false
		if _, ok := set[r]; ok {
			return false
			// иначе это уникальный символ добавляем в мапу
		} else {
			set[r] = struct{}{}
		}
	}

	// если все символы добавлены и функция не вышла с false, то вернём true
	return true
}

// основной поток
func main() {

	// входной массив
	str := []string{"abcd", "abCdefAaf", "aabcd"}

	// проходимся по всему массиву
	for _, s := range str {
		fmt.Printf("%s - %v\n", s, uniqueChars(s))
	}

}
