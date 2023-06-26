package main

import (
	"time"
)

// с помощью канала
func mySleep(duration time.Duration) {
	// в канал поступит данные через duration
	// он прочитается и горутина разблокируется
	<-time.After(duration)
}

// основной поток
func main() {
	// вызов моего слипера
	mySleep(10 * time.Second)
}
