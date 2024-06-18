package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// Т.к. в данном задании горутины будут выполнять CPU-bound работу, то не имеет смысла создавать горутин более чем
	// GOMAXPROCS, т.к. при большем количестве горутин, будет тратиться дополнительное время на переключение контекста,
	// если горутин будет <= GOMAXPROCS, то они смогут работать параллельно и без переключения контекста
	arr := [5]int{2, 4, 6, 8, 10}
	var numThreads int
	if maxProcs := runtime.GOMAXPROCS(0); len(arr) > maxProcs {
		numThreads = maxProcs
	} else {
		numThreads = len(arr)
	}
	chunk := len(arr) / numThreads
	var sum int64
	var wg sync.WaitGroup

	// Суммируем квадраты элементов массива, lock-free методом, т.к. он это будет работать быстрее чем при использовании мьютекса
	// Для каждой горутины определяем срез в котором она будет считать квадраты элементов, в случае размера
	// массива <= GOMAXPROCS, каждая горутина будет считать квадрат одного элемента
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		start := i * chunk
		end := start + chunk
		if i == numThreads-1 {
			end = len(arr)
		}
		go func() {
			defer wg.Done()
			for _, elem := range arr[start:end] {
				sqr := int64(elem * elem)
				atomic.AddInt64(&sum, sqr)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Sum:", sum)
}
