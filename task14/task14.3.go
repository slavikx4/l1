package main

import (
	"fmt"
)

// основной поток
func main() {

	// массив входных данных
	array := []any{10, "dog", true, make(chan int)}

	// проходимся по каждому элементу массива
	for _, el := range array {
		// определяем тип с помощью управляющих символов
		fmt.Printf("%v\t%T\n", el, el)
	}
}
