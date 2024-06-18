package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 50)
	b := big.NewInt(3 << 27)

	// Умножение
	mul := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение: %v * %v = %v\n", a, b, mul)

	// Деление
	div := new(big.Int).Div(a, b)
	fmt.Printf("Деление: %v / %v = %v\n", a, b, div)

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сложение: %v + %v = %v\n", a, b, sum)

	// Вычитание
	sub := new(big.Int).Sub(a, b)
	fmt.Printf("Вычитание: %v - %v = %v\n", a, b, sub)
}
