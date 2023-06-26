package main

import "fmt"

// основной поток
func main() {

	// массив входных данных
	array := []any{10, "dog", true, make(chan int)}

	// вызов функции опредеяющая тип
	printType(array...)
}

// функция опредеяющая тип данных
func printType(i ...any) {
	// проходимся по всем входным данным
	for _, el := range i {
		// специальная конструкция для определения типа
		switch el.(type) {
		case int:
			fmt.Printf("%v\tit`s int\n", el)
		case string:
			fmt.Printf("%v\tit`s string\n", el)
		case bool:
			fmt.Printf("%v\tit`s bool\n", el)
		// таким способ можно определить конкретный канал, но нету просто канала chan как тип
		default:
			fmt.Printf("%v\tunknown\n", el)
		}
	}
}
