package main

import (
	"fmt"
	"math/rand"
)

// основной поток
func main() {

	// два множества с неопределённым типом данных
	array1 := make([]any, 20)
	array2 := make([]any, 20)

	// инициализация множеств
	for i := 0; i < 20; i++ {
		array1[i] = rand.Intn(20)
		array2[i] = rand.Intn(20)
	}

	// вызов функции для нахождения пересечения
	res := intersection(array1, array2)
	// вывод результата
	fmt.Println(res)
}

// intersection функция, которая находит пересечение между двумя множествами
func intersection(array1, array2 []any) []any {

	// временная мапа для реализации механизма поиска пересечений
	temp := make(map[any]struct{})

	// вместо первого массива создаём мапу, где нам не важно что будет лежать в значении для ключа
	for i := 0; i < len(array1); i++ {
		temp[array1[i]] = struct{}{}
	}

	// очищаем первый массив, потому что дальше здесь будет храниться результат
	array1 = []any{}

	// перебираем элементы второго массива
	for _, el := range array2 {
		// если в мапе существует ключ, то значит это пересечение
		if _, ok := temp[el]; ok {
			// записываем в результативный массив
			array1 = append(array1, el)
		}
	}

	// возвращаем массив с результатом
	return array1
}
