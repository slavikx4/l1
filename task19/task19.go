package main

// основной поток
func main() {

	// входная строка
	str := "asd123фыв"

	// слайс рун из строки, чтобы корректно работать с юникодом
	res := []rune(str)

	// длина слайса рун
	l := len(res)

	// проходим до середины слайса
	for i := 0; i < l/2; i++ {
		// свапаем элементы слайса  левый с правым
		res[i], res[l-1-i] = res[l-1-i], res[i]
	}
}
