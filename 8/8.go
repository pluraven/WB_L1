package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func main() {
	var num int64
	var i, value byte
	fmt.Print("Введите число в десятичном виде: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Число в двоичном виде до операции:", strconv.FormatInt(num, 2))
	fmt.Print("Введите номер бита: ")
	_, err = fmt.Scan(&i)
	if err != nil {
		log.Fatal(err)
	}
	// в int64 числе можно изменить только 63 бита, т.к. 64 является знаковым
	if i > 62 {
		log.Fatal(errors.New("Слишком большая позиция для изменения. Выберете число из диапазона [0, 62]"))
	}
	fmt.Print("Введите новое значение бита: ")
	_, err = fmt.Scan(&value)
	if err != nil {
		log.Fatal(err)
	}

	switch value {
	case 0:
		num &= ^(1 << i)
	case 1:
		num |= 1 << i
	default:
		log.Fatal(errors.New("Неверное значение. Бит может принимать значения 0 или 1."))
	}
	fmt.Println("Число после операции:", strconv.FormatInt(num, 2))
}
