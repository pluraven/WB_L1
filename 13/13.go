package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	a := int32(10)
	b := int32(5)
	fmt.Println("a:", a, "b:", b)
	swap1(&a, &b)
	fmt.Println("Swapped.")
	fmt.Println("a:", a, "b:", b)
	swap2(&a, &b)
	fmt.Println("Swapped.")
	fmt.Println("a:", a, "b:", b)
	swap3(&a, &b)
	fmt.Println("Swapped.")
	fmt.Println("a:", a, "b:", b)
}

// Существует два способа обменять значения переменных без использования временной переменной

// 1. Способ - синтаксический сахар языка Go, - под капотом используются атомики
func swap1(a, b *int32) {
	*a, *b = *b, *a
}

// 2. Способ использовать атомарные операции напрямую
func swap2(a, b *int32) {
	// Используем атомики для безопасного обмена значений
	atomic.StoreInt32(a, atomic.SwapInt32(b, atomic.LoadInt32(a)))
}

// 3. Способ использовать битовые операции, сохранять те биты которые различаются в двух числах
func swap3(a, b *int32) {
	*b = *b ^ *a
	*a = *b ^ *a
	*b = *b ^ *a
}
