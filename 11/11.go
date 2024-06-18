package main

import (
	"fmt"
)

func main() {
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}
	set2 := map[int]bool{1: true, 3: true, 5: true, 183: true, 11: true}

	intersectionSet := intersection(set1, set2)

	fmt.Println("Пересечение множеств:")
	fmt.Println(intersectionSet)
}

func intersection(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)
	for key := range set1 {
		if set2[key] {
			result[key] = true
		}
	}
	return result
}
