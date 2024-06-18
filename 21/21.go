package main

import "fmt"

// Интерфейс Printer, который ожидает клиент
type Printer interface {
	PrintName()
}

// Структура Person с несовместимым интерфейсом
type Person struct {
	name string
}

func (p *Person) ShowName() {
	fmt.Println(p.name)
}

// Адаптер - реализует интерфейс Printer
type PersonAdapter struct {
	person *Person
}

// Метод PrintName для адаптера PersonAdapter
func (p *PersonAdapter) PrintName() {
	p.person.ShowName()
}

func main() {
	Bob := &Person{name: "Bob"}
	// Создаем адаптер, который реализует интерфейс Printer
	var adapter Printer = &PersonAdapter{
		person: Bob,
	}
	// Используем метод интерфейса
	adapter.PrintName()
}
