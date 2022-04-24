package main

import (
	"fmt"
	"go_for_algorithm/src/search"
)

func main() {
	for value := 1; value <= 10; value++ {
		fmt.Println(search.BinarySearch([]int{}, value))
	}
}
