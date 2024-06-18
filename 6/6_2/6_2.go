package main

import (
	"fmt"
	"sync"
	"time"
)

// Останавливаем горутину через закрытие канала

func main() {
	var wg sync.WaitGroup
	stopChannel := make(chan bool)
	wg.Add(1)

	go Worker(stopChannel, &wg)
	time.Sleep(4 * time.Second)

	close(stopChannel)
	wg.Wait()
}

func Worker(stopChannel chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Горутина завершена.")
	// Имитируем полезную работу
	fmt.Println("Горутина работает...")
	_, ok := <-stopChannel
	if !ok {
		return
	}
}

// Также при необходимости получения данных из канала аналогичный результат можно получить используя range по каналу
func Worker2(stopChannel chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Горутина завершена.")
	for _ = range stopChannel {
		// Имитируем полезную работу (получаем данные)
		fmt.Println("Горутина работает...")
	}
}
