package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// основной поток
func main() {

	// массив входных данных
	array := make([]int, 100)

	// инициализируем массив
	for i := 0; i < 100; i++ {
		array[i] = rand.Intn(100)
	}

	// сортируем массив
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	// поиск индекса числа
	res := binarySearch(array, 10)
	fmt.Println(res)
}

// функция бинарного поиска
func binarySearch(array []int, num int) int {

	// определяем границы поиска
	left, right := 0, len(array)

	// пока не искомая область нестанет равняться 1
	for left < right {
		pivot := (left + right) / 2
		// если пивот подошёл, то возвращаем индес - ответ
		if array[pivot] == num {
			return pivot
			// иначе если этот элемент больше, чем число, то следует искать где-то левее от пивота
		} else if array[pivot] > num {
			right = pivot - 1
			// иначе если под данным пивотом лежит число меньше искомого, то искать нужно в правой части
		} else {
			left = pivot + 1 // если меньше, то на 1 больше
		}
	}

	// если такого элемента нет
	return -1
}
