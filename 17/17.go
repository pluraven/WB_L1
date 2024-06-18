package main

import "fmt"

func BinarySearch(arr []int, el int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == el {
			return mid
		} else if arr[mid] < el {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	index := BinarySearch(arr, 7)
	fmt.Println("Индекс элемента 7:", index)
}
