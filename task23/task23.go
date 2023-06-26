package main

// основной поток
func main() {

	// входной массив
	array := []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// массив с удалённым индексом
	array = deleteElem(array, 9)
}

// функция удаления элемента по индексу
func deleteElem(array []any, i int) []any {
	// делаем новый слайс
	res := make([]any, len(array)-1)

	// копируем значения из страого слайса в новый
	copy(res[:i], array[:i])
	copy(res[i:], array[i+1:])

	// возвращаем результат
	return res
}
