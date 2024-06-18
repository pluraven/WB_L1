package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	const numberOfWorkers = 3
	// Используем RWMutex для безопасного доступа к map и для получения возможности чтения из map во время записи
	var mu sync.RWMutex
	var wg sync.WaitGroup
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	wg.Add(numberOfWorkers)
	m := make(map[int]bool)

	for i := 0; i < numberOfWorkers; i++ {
		go func(ctx context.Context, wg *sync.WaitGroup) {
			defer wg.Done()
			defer fmt.Println("Горутина уничтожена.")
			fmt.Println("Создана горутина.")
			for {
				select {
				case <-ctx.Done():
					return
				default:
					key := rand.Intn(10000)
					mu.Lock()
					m[key] = true
					mu.Unlock()
					fmt.Println("В map записано новое значение.")
				}
			}
		}(ctx, &wg)
	}
	wg.Wait()
}
