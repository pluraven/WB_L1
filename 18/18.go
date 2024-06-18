package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Данную программу можно реализовать с применением Mutex, однако атомики работают быстрее

// Структура Counter - представляет счетчик
type Counter struct {
	count int64
}

// Метод для инкрементирования счетчика
func (c *Counter) Increment() {
	atomic.AddInt64(&c.count, 1)
}

// Функция для получения текущего значения счетчика
func (c *Counter) Get() int64 {
	return c.count
}

func main() {
	const (
		numWorkers = 10
		iterations = 1000
	)
	counter := Counter{}
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Ожидаемое значение счетчика:", numWorkers*iterations)
	fmt.Println("Итоговое значение счетчика:", counter.Get())
}
