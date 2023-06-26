package main

import (
	"fmt"
	"sync"
)

// способ с использованием waitGroup

// основной поток
func main() {
	// waitGroup для ожидания завершения всех горутин
	wg := &sync.WaitGroup{}

	numbers := [...]int{2, 4, 6, 8, 10}

	// проходимся по массиву чисел и передаём элемент в новую горутину на обработку
	for _, number := range numbers {
		// добавление к счётчику работающих горутин 1
		wg.Add(1)
		//запуск нового потока
		go squaring(wg, number)
	}

	//ожидает пока счётчик работающих горутин не станет равен 0
	wg.Wait()

	//вывод чисел будет в рандомном порядке
}

// принимает указатель на waitGroup, чтобы не была создана копия
func squaring(wg *sync.WaitGroup, number int) {
	//у fmt.Printf под капотом вызывается fmt.Fprintf, которая и записывает в os.Stdout
	fmt.Printf("%d\t", number*number)
	// уменьшение счётчика запущенных горутин на 1
	wg.Done()
}
