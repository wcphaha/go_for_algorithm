package main

import (
	"fmt"
	"go_for_algorithm/src/search"
)

func main() {
	for value := 1; value <= 10; value++ {
		fmt.Println(value)
		fmt.Println(search.TraversingSearch([]int{1, 2, 3, 4, 4, 5, 6, 6, 7, 8}, value))
	}
}
