package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	// Группа ожидания для того чтобы дождаться завершения всех горутин
	var wg sync.WaitGroup
	fmt.Print("Array: ")
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Print(arr[i], " ")
		}()
	}
	wg.Wait()
	fmt.Println()
}
