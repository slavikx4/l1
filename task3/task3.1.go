package main

import (
	"fmt"
	"sync"
)

// способ с использованием waitGroup и Mutex

// основноцй поток
func main() {

	// waitGroup для ожидания завершения всех необходимых вычислений
	wg := &sync.WaitGroup{}

	// создаём наш counter
	count := &counter{
		summa: 0,
		mu:    sync.Mutex{},
	}

	numbers := [...]int{2, 4, 6, 8, 10}

	// передаём числа в горутины
	for _, number := range numbers {
		// добавляем  к счётчику незавершённых горутин единицу
		wg.Add(1)
		//запускаем горутину с очередным числом
		go squaring(wg, count, number)
	}

	// дожидаемся, пока все необходимые вычисления выполнятся и счётчик станет равен 0
	wg.Wait()

	// выводим сумму
	fmt.Println("summa: ", count.summa)
}

// counter нужен для исключения гонки данных - добавление мьютексов при работе с обычным int
type counter struct {
	summa int
	mu    sync.Mutex
}

// squaring возводит и прибавляет к общей сумме очерендное поступившее число
func squaring(wg *sync.WaitGroup, count *counter, number int) {
	// блокируем доступ для остальных горутин
	count.mu.Lock()
	// добавляем к общей сумме квадрат числа
	count.summa += number * number
	// освобождаем доступ для остальных горутин
	count.mu.Unlock()

	// убавляем счтётчик на единицу
	wg.Done()
}
