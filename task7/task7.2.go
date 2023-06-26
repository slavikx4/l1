package main

import (
	"sync"
	"time"
)

// способ c обычной мапой и мьютексом в структуре

type myMap struct {
	mp map[int]int
	mu sync.Mutex
}

// метод добавления в нашу мапу значений
func (m *myMap) add(key, value int) {
	m.mu.Lock()
	m.mp[key] = value
	m.mu.Unlock()
}

// основной поток
func main() {

	// создаём объект нашей мапы
	myMp := &myMap{
		mp: make(map[int]int),
		mu: sync.Mutex{},
	}

	// запускаем горутины
	for i := 0; i < 100; i++ {
		go myMp.add(i, i+100)
	}

	// ожидаем пока все горутины выполнятся. не используем  waitGroup, так как
	// в реальной практике нам не нужно ждать пока значения запишутся
	// Sleep вызван только для демонстрации работоспособности решения
	time.Sleep(10 * time.Second)
}
