package main

import (
	"fmt"
	"math"
)

// структура точки
type Point struct {
	x int
	y int
}

// конструктор
func NewPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// геттер для поля x
func (p *Point) GetX() int {
	return p.x
}

// геттер для поля y
func (p *Point) GetY() int {
	return p.y
}

// основной поток
func main() {

	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)

	res := delta(p1, p2)
	fmt.Println(res)
}

// возвращает расстояние между точками a, b *Point
func delta(p1, p2 *Point) float64 {
	x := p1.GetX() - p2.GetX()
	y := p1.GetY() - p2.GetY()
	// по т Пифагора
	return math.Sqrt(float64(x*x + y*y))
}
