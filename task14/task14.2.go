package main

import (
	"fmt"
	"reflect"
)

// основной поток
func main() {

	// массив входных данных
	array := []any{10, "dog", true, make(chan int)}

	// проходим по всему массиву входных данных
	for _, el := range array {
		// с помощью reflect определяем тип
		fmt.Printf("%v\t%v\n", el, reflect.TypeOf(el))
	}
}
