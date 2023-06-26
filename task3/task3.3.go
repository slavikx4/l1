package main

import "fmt"

// с использованием каналов без магический чисел

// основной поток
func main() {

	// переменная с информацией о сумме
	var summa int

	numbers := [...]int{2, 4, 6, 8, 10}

	// канал для передачи ещё не обработаных чисел
	inputChan := make(chan int)

	// канал для передачи уже возведённых в квадрат чисел
	outputChan := make(chan int)

	// новый поток для возведение чисел в квадрат
	// в конце закрываем канал outputChan, чтобы сообщить, что входящих данных по этому каналу больше не ожидается
	go func() {
		defer close(outputChan)
		// считываем данные из канала и обрабатываем
		for num := range inputChan {
			outputChan <- num * num
		}
	}()

	// новый поток для передачи входных данных
	// в конце закрываем канал inputChan, чтобы сообщить, что входящих данных по этому каналу больше не ожидается
	// необходимо вызвать функцию в новом потоке, чтобы не произошёл deadlock:
	// передавая значение в inputChan, тем самым записываем в outputChan после вычислений,
	// а из outputChan не считываются данные => следующий раз нельзя записать в outputChan ,  и тут вылезает deadlock
	go func() {
		defer close(inputChan)
		// передаём данные в канал для дальнешего возведения в квадрат
		for _, num := range numbers {
			inputChan <- num
		}
	}()

	// стопим основной поток, пока не закроется канал outputChan
	// суммируем поступившие данные
	for out := range outputChan {
		summa += out
	}

	// выводим сумму
	fmt.Println(summa)
}
