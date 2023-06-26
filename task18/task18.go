package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// структура счётчик, с мьютексом для ликвидации гонки данных
type counter struct {
	summa int
	mu    sync.Mutex
}

// основной поток
func main() {

	// создаём объект нашего счётчика
	count := &counter{
		summa: 0,
		mu:    sync.Mutex{},
	}

	// второй способ - атомарный счётчик
	var count2 int32

	// запускаем миллион горутин
	for i := 0; i < 1000000; i++ {
		go func() {
			// первый способ
			count.mu.Lock()
			count.summa++
			count.mu.Unlock()

			// второй способ
			atomic.AddInt32(&count2, 1)
		}()
	}

	// ждём пару секунд, чтоы все горутины выполнились( можно было и waitGroup поставить)
	time.Sleep(5 * time.Second)

	// выводим результат
	// 1 способ
	fmt.Println(count.summa)
	// 2 способ
	fmt.Println(count2)
}
