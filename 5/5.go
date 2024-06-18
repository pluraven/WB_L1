package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	defer fmt.Println("Завершение работы...")
	var n int
	var wg sync.WaitGroup
	fmt.Print("Введите количество секунд до завершения программы: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan int)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	wg.Add(2)
	go Receiver(ctx, ch, &wg)
	go Sender(ctx, ch, &wg)
	<-ctx.Done()
	wg.Wait()
}

func Sender(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- rand.Intn(100):
			fmt.Println("Успешно отправлено.")
		}
	}
}

func Receiver(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ch:
			fmt.Println("Успешно прочитано.")
		}
	}
}
