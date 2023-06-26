package main

import "fmt"

// способ с использованием каналов с "магическими числами"

// основной канал
func main() {

	// заводим переменную для хранения суммы
	var summa int

	numbers := [...]int{2, 4, 6, 8, 10}

	// создаём небуфиризированный канал
	chanInput := make(chan int)

	// запускаем безымянную горутину
	go func() {
		// закрываем канал, хоть здесь закрытие канала и не играет роли
		defer close(chanInput)
		// так как мы знаем что в массиве 5 чисел
		for i := 0; i < 5; i++ {
			// читаем из канала
			num := <-chanInput
			// возводим в квадрат и сумируем с общей суммой
			summa += num * num
		}
	}()

	// передаём числа в канал
	for _, num := range numbers {
		chanInput <- num
	}

	// выводим сумму
	fmt.Println(summa)
}
