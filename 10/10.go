package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int64][]float64)
	for _, elem := range temperatures {
		key := int64(elem) / 10 * 10
		if _, ok := groups[key]; !ok {
			groups[key] = make([]float64, 0)
		}
		group := groups[key]
		group = append(group, elem)
		groups[key] = group
	}
	fmt.Println(groups)
}
