package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var numWorkers int
	fmt.Print("Введите желаемое количество воркеров: ")
	_, err := fmt.Scan(&numWorkers)
	if err != nil {
		log.Fatal(err)
	}
	// В данной ситуации уместно использовать context для завершения работы воркеров, т.к. context позволяет удобно
	// управлять жизненным циклом воркеров.
	ctx, cancel := context.WithCancel(context.Background())
	// Канал интерфейсного типа для того, чтобы было возможно записывать и считывать из канала данные любого типа,
	// как указано в задании
	ch := make(chan interface{})
	// Отслеживаем SIGINT
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	for i := 0; i < numWorkers; i++ {
		go RunWorker(ctx, ch)
	}
	// Постоянно записываем данные в канал в главном потоке
	for {
		select {
		case <-sig:
			cancel()
			fmt.Println("Вывод завершён.")
			return
		default:
			ch <- rand.Intn(100)
		}
	}
}

func RunWorker(ctx context.Context, ch chan interface{}) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(<-ch)
		}
	}
}
