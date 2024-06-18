package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = make(chan int)
	var b interface{} = make(chan map[int]string)
	fmt.Println("Data type:", checkTypeWithSwitch(a))
	fmt.Println("Data type:", checkTypeWithReflect(b))
}

// Можно использовать встроенную возможность определения типа переменной с конструкцией switch, однако возникает
// проблема с определением канала как типа
func checkTypeWithSwitch(elem interface{}) string {
	switch elem.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	case chan bool:
		return "chan bool"
	default:
		return "unknown"
	}
}

// Также можно использовать рефлексию
func checkTypeWithReflect(elem interface{}) string {
	typ := reflect.TypeOf(elem)
	return fmt.Sprintf("%v", typ)
}
