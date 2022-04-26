package main

import (
	"fmt"
	"go_for_algorithm/src/sort"
)

func main() {
	nums := []int{1, 3, 34, 123, 34, 12, 456, 12, 4}
	heapSort := sort.Heap{}
	fmt.Println(heapSort.Sort(nums, true))
}
