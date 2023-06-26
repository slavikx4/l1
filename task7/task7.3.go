package main

import (
	"sync"
	"time"
)

// способ с использованием sync.Map

// основной поток
func main() {

	// данная карта сама обеспечивает безопастность, без дополнительной связки с мьютексом
	mp := &sync.Map{}

	// запускаем горутины , вызывая метод у мапы
	for i := 0; i < 1000; i++ {
		go mp.Store(i, i+100)
	}

	// ожидаем пока все горутины выполнятся. не используем  waitGroup, так как
	// в реальной практике нам не нужно ждать пока значения запишутся
	// Sleep вызван только для демонстрации работоспособности решения
	time.Sleep(10 * time.Second)
}