package main

import (
	"math"
)

// createStep функция ищет промежуток для температуры
func createStep(f float64) int {
	return int(math.Floor(f/10)) * 10
}

// основной поток
func main() {

	// слайс с входными данными
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// карта для хранения уже обработаной информации
	mp := make(map[int][]float64)

	// перебираем каждую температуру
	for _, temperature := range temperatures {
		// вызывая функцию, находим промежуток
		step := createStep(temperature)
		// определяем встречался ли раньше такой промежуток
		_, ok := mp[step]
		// если не встречался, то создаём новый слайс
		if !ok {
			mp[step] = []float64{temperature}
			// иначе добавляем в уже существующий слайс
		} else {
			mp[step] = append(mp[step], temperature)
		}
	}
}
