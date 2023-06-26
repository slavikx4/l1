package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os/signal"
	"syscall"
	"time"
)

// main основной поток
func main() {

	var counterWorkers int

	// при помощи флага при старте программы указываем колличество воркеров
	flag.IntVar(&counterWorkers, "counterWorkers", 10, "enter counter workers")
	flag.Parse()

	// создаём небуфиризированный канал для передачи данных в воркеры
	mainChan := make(chan int)

	// создаём контекст, который будет завершаться при подаче сингнала, вызванного нажатиес Ctrl + C
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT)

	// даём старт воркерам, создав counterWorkers горутин
	for i := 0; i < counterWorkers; i++ {
		go func() {
			// пока канал не закрыт, свободный воркер будет брать данные из основного канала
			for out := range mainChan {
				fmt.Printf("%d\t", out)
			}
		}()
	}

	// метка loop позволяет вызвать break у определённого итератора
loop:
	// запускаем бесконечный итератор
	for {
		// спим 1 секунду, чтобы понимать, что выводится в консоль
		// (да, в таком случаем не нужно много воркеров, так как и один воркер будет справляться,
		// но зато в консоли не творится хаус)
		time.Sleep(1 * time.Second)
		select {
		//если поступил сигнал, то контекст можно отменить
		case <-ctx.Done():
			// закрываем канал, чтобы завершилась работа горутин(воркеров)
			close(mainChan)
			fmt.Println("EXIT")
			//завершаем работу итератора, чтобы выйти из программы
			break loop
		// если контекст не поступил, то выполняем дефолтную интрукцию
		default:
			// записываем рандомное int число в канал
			mainChan <- rand.Int()
		}
	}
}
