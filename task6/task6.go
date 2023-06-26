package main

func main() {

	// 0 способ
	go gorutine0()

	// 1 способ
	go gorutine1()

	// 2 способ
	ch2 := make(chan int)
	go gorutine2(ch2)
	close(ch2)

	// 3 способ
	ch3 := make(chan struct{})
	go gorutine3(ch3)
	ch3 <- struct{}{}

	// 4 способ
	ch4 := make(chan struct{})
	go gorutine4(ch4)
	close(ch4)

}

// закончив выполнять работу
func gorutine0() {
	// вычисления
}

// с помощью return
func gorutine1() {
	//различные вычисления, может быть проверка условий
	return
}

// с помощью закрытия канала
func gorutine2(ch chan int) {
	for c := range ch {
		c++
	}
}

// с помощью передачи в канал и for-select (сюда же отмена с контекст)
func gorutine3(ch chan struct{}) {
	for {
		select {
		case <-ch:
			return
		default:
			// вычисления
		}
	}
}

// при закрытии канала и for-select
func gorutine4(ch chan struct{}) {
	for {
		select {
		case _, ok := <-ch:
			// проверка: если канал закрыт, то выходим
			if !ok {
				return
			}
			//вычисления
		default:
			// вычисления
		}
	}
}
