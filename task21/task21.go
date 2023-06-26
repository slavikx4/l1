package main

import "fmt"

// представим, что есть два компьютера: mac и windows.
// на mac квадратный usb порт, а на windows круглый usb порт
// но у нас есть клиент, который хочет вставить именно квадратный usb провод

// интерфейс компьютера
type computer interface {
	insertSquareUsbInComputer()
}

// структура компьютера с квадратным портом
type mac struct{}

// метод mac реализующий интерфейс computer
func (m *mac) insertSquareUsbInComputer() {
	fmt.Println("в мак вставлен квадратный usb провод")
}

// структура компьютера с круглым usb портом
type windows struct{}

// метод windows который подразумевает вставку круглого usb
func (w *windows) insertCircleUsbInComputer() {
	fmt.Println("в виндувс вставлен круглый usb провод")
}

// структура адаптер для windows
type windowsAdapter struct {
	windows *windows
}

// метод адаптера, чтобы реализующий интерфейс computer
func (w *windowsAdapter) insertSquareUsbInComputer() {
	w.windows.insertCircleUsbInComputer()
}

// структура клиента
type client struct{}

// метод client подразумевающий вставку квадратного usb провода в компьютер
func (c *client) insertSquareUsbInComputer(com computer) {
	com.insertSquareUsbInComputer()
}

// основной поток
func main() {

	// объект клиента
	client := &client{}

	// объект мака
	mac := &mac{}

	// клиент вызывает метод, передав аргументом объект реализующий интерфейс computer
	client.insertSquareUsbInComputer(mac)

	// объект виндус
	windows := &windows{}

	// объект адаптера виндуса, принимающий в конструктор объект виндус
	windowsAdapter := &windowsAdapter{windows: windows}

	// клиент вызывает метод, передав аргументом объект реализующий интерфейса computer
	// но на самом деле метод будет вызван у объекта виндус внутри
	client.insertSquareUsbInComputer(windowsAdapter)
}
