package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup

	arr := [...]int{1, 2, 3, 4, 5}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, x := range arr {
			ch1 <- x
		}
		close(ch1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for x := range ch1 {
			ch2 <- x * 2
		}
		close(ch2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range ch2 {
			fmt.Println(result)
		}
	}()

	wg.Wait()
}
