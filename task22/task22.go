package main

import (
	"fmt"
	"math/big"
)

// основной поток
func main() {

	// создаём три переменных для хранения больших чисел
	var a, b, c = &big.Int{}, &big.Int{}, &big.Int{}

	// инициализируем две переменный, а третью не трогаем
	a.Exp(big.NewInt(2), big.NewInt(22), nil)
	b.Exp(big.NewInt(2), big.NewInt(21), nil)

	// в третей переменной будет храниться результат вычислений первых двух
	c.Add(a, b)
	fmt.Printf("сумма: %v\n", c.String())
	c.Mul(a, b)
	fmt.Printf("умножение: %v\n", c.String())
	c.Sub(a, b)
	fmt.Printf("вычитание: %v\n", c.String())
	c.Div(a, b)
	fmt.Printf("деление: %v\n", c.String())
}
