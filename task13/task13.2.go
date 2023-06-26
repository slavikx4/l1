package main

// основной поток
func main() {

	// инициализируем два числа
	first, second := 1, 2

	// несложными математическими действиями  меняем значения
	first += second
	second = first - second
	first -= second
}
