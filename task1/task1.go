package main

import "fmt"

// Human структура родитель
type Human struct {
	name string
	age  uint
}

// happyBirthday метод структуры Human
func (h *Human) happyBirthday() {
	h.age++
	fmt.Println("[h] current age: ", h.age)
}

// printName метод структуры Human
func (h *Human) printName() {
	fmt.Println("[h] name: ", h.name)
}

// Action дочерняя структура
type Action struct {
	Human      // вложенная структура Human
	age   uint // такое же поле, что и в Human
}

// happyBirthday метод структуры Action, переопределяющий поведение Human
func (a *Action) happyBirthday() {
	a.age++
	fmt.Println("[a] current age: ", a.age)
}

// setName метод Action, перезаписывающий поле у Human
func (a *Action) setName(name string) {
	a.name = name
}

// основной поток
func main() {
	h := Human{
		name: "slava",
		age:  15,
	}
	h.happyBirthday() // [h] current age: 16

	//передаём в параметры возраст Action и родителя
	a := Action{age: 10, Human: h}

	// вызовится переписаный метод Action
	a.happyBirthday() // [a] current age: 11
	// явно указали, что  метод нужно вызвать у Human, который лежит внутри Action
	a.Human.happyBirthday() // [h] current age: 17
	// продемонстратировано, что поле age у Action и Human разные
	a.happyBirthday() // [a] current age: 12

	// к полю name можно обратиться через Action напрямую
	a.printName() // [h] name: slava

	// продемонстратировано, что можно изменять поля родителя, через дочерний класс
	a.setName("artem")
	a.printName() // [h] name: artem
}
