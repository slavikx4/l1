package main

import (
	"strings"
)

// основной поток
func main() {

	// входная строка
	str := "snow dog sun"

	// разбиваем строку на слайс
	array := strings.Fields(str)

	// находим длину слайса
	l := len(array)

	// проходим до середины слайса
	for i := 0; i < l/2; i++ {
		// свапаем левый и правый элемент
		array[i], array[l-1-i] = array[l-1-i], array[i]
	}
}
