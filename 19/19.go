package main

import "fmt"

func main() {
	str := "sаффа"
	fmt.Println(Reverse(str))
}

func Reverse(str string) string {
	buff := []rune(str)
	result := make([]rune, len(buff))
	for i := len(buff) - 1; i >= 0; i-- {
		result[len(buff)-1-i] = buff[i]
	}
	return string(result)
}
