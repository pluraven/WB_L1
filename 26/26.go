package main

import (
	"fmt"
	"unicode"
)

func main() {
	str1 := "Нн"
	fmt.Println(str1, "-", isAllUnique(str1))
	str2 := "qwerty"
	fmt.Println(str2, "-", isAllUnique(str2))
	str3 := "ОO" // один символ ascii, другой unicode из русской раскладки
	fmt.Println(str3, "-", isAllUnique(str3))
}
func isAllUnique(str string) bool {
	tmpMap := make(map[rune]bool)
	buff := []rune(str)
	for _, el := range buff {
		tmpMap[unicode.ToLower(el)] = true
	}
	if len(tmpMap) == len(buff) {
		return true
	}
	return false
}
