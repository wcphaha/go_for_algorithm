package sort

import (
	"fmt"
)

func ExampleHeap_Sort() {
	nums := []int{1, 3, 34, 123, 34, 12, 456, 12, 4}
	heapSort := Heap{}
	fmt.Println(heapSort.Sort(nums, true))
}
