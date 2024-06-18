package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)
	for _, key := range strings {
		set[key] = true
	}

	fmt.Println("Множество:")
	fmt.Println(set)
}
