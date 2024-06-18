package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "sнow дog sun"
	fmt.Println("Строка до операции:", str)
	fmt.Println("Строка после операции:", ReverseWordsInString(str))
}

func ReverseWordsInString(str string) string {
	words := strings.Fields(str)
	result := make([]string, len(words))
	for i := 0; i < len(words); i++ {
		result[len(words)-1-i] = words[i]
	}
	return strings.Join(result, " ")
}
