package main

import (
	"fmt"
	"math/rand"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// При данном подходе в памяти останется большая строка, которая будет занимать место на диске, т.к. на нее будет
	// ссылаться justString
	// justString := v[:100]
	// Решением является создание новой строки на определенное количество символов, GC отчистит память,
	// занимаемою большой строкой
	justString = string(CopyString([]rune(v), 100))
}

func main() {
	someFunc()
	fmt.Println(justString)
}

func createHugeString(size int) string {
	result := make([]rune, size)
	letters := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 'q', 'w', 'e', 'r', 't', 'ф'}
	for i := 0; i < size; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func CopyString(s []rune, size int) []rune {
	result := make([]rune, size)
	for i := 0; i < size; i++ {
		result[i] = s[i]
	}
	return result
}
