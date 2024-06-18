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
	// Используем sync.Map - работает быстрее, чем обычный map с Mutex
	var m sync.Map
	var wg sync.WaitGroup
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	wg.Add(numberOfWorkers)

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
					m.Store(key, true)
					fmt.Println("В map записано новое значение.")
				}
			}
		}(ctx, &wg)
	}
	wg.Wait()
}
