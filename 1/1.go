package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Action struct {
	Human
}

func main() {
	Bob := Action{Human{name: "Bob", age: 25}}
	// Вызов метода принадлежащего родительской структуре у экземляра струтуры Action
	Bob.Walk()
}

func (h *Human) Walk() {
	fmt.Println(h.name, "is walking.")
}

func (h *Human) Sleep() {
	fmt.Println(h.name, "is sleeping.")
}
