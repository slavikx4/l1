package main

import "fmt"

// основной поток
func main() {

	// массив входных данных
	array := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// канал для входа
	chanInput := make(chan int)
	// канал для выхода
	chanOutput := make(chan int)

	// новый поток обрабатывающий входной поток
	go func() {
		defer close(chanOutput)
		for num := range chanInput {
			chanOutput <- num * 2
		}
	}()

	// новый поток записывающий во входной поток
	go func() {
		defer close(chanInput)
		for _, num := range array {
			chanInput <- num
		}
	}()

	// чтение из выходного потока
	for out := range chanOutput {
		fmt.Printf("%d\t", out)
	}
}
