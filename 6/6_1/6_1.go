package main

import (
	"fmt"
	"sync"
	"time"
)

// Отправляем стоп сигнал через канал

func main() {
	var wg sync.WaitGroup
	stopChannel := make(chan bool)
	wg.Add(1)

	go Worker(stopChannel, &wg)
	time.Sleep(4 * time.Second)

	stopChannel <- true
	wg.Wait()
}

func Worker(stopChannel chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Горутина завершена.")
	// Имитируем полезную работу
	for {
		select {
		case <-stopChannel:
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Горутина работает...")
		}
	}
}
