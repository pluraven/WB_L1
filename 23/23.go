package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("arr до удаления элементов: ", arr)
	arr = removeElementUsingAppend(arr, 3)
	fmt.Println("arr после удаления 3 элемента: ", arr)
	arr = removeElementUsingCopy(arr, 3)
	fmt.Println("arr после удаления 3 элемента: ", arr)
}

func removeElementUsingAppend(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func removeElementUsingCopy(slice []int, index int) []int {
	result := make([]int, len(slice)-1)
	copy(result, slice[:index])
	copy(result[index:], slice[index+1:])
	return result
}
