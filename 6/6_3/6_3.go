package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Останавливаем горутину с помощью контекста

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)

	go Worker(ctx, &wg)
	time.Sleep(4 * time.Second)

	cancel()
	wg.Wait()
}

func Worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Горутина завершена.")
	// Имитируем полезную работу
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Горутина работает...")
		}
	}
}
