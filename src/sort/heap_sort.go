package sort

/*Heap
@Description: the heap class with sort method
*/
type Heap struct {
	kind   string
	nums   []int
	length int
}

/*
	@Title Sort
 	@Description: heap-sort
 	@Author WuChengPei 2022-04-26 15:15:35
 	@Param nums:
 	@Param reverse: sort order
 	@Return res: sorted sequence
*/
func (heap *Heap) Sort(nums []int, reverse bool) (res []int) {
	heap.BuildHeap(nums, reverse)
	for heap.length > 0 {
		res = append(res, heap.nums[0])
		heap.nums = append(heap.nums[:0], heap.nums[1:]...)
		heap.length--
		heap.BuildHeap(heap.nums, reverse)
	}
	return res
}

/*
	@Title BuildHeap
 	@Description: build a heap
 	@Author WuChengPei 2022-04-26 15:17:41
 	@Param nums:
 	@Param reverse:
*/
func (heap *Heap) BuildHeap(nums []int, reverse bool) {
	if reverse {
		heap.kind = "MIN"
	} else {
		heap.kind = "MAX"
	}
	heap.nums = nums
	heap.length = len(nums)
	for i := (heap.length - 2) / 2; i >= 0; i-- {
		heap.Heapify(i)
	}

}

func (heap *Heap) SwapNode(i, j int) {
	heap.nums[i], heap.nums[j] = heap.nums[j], heap.nums[i]
}

/*
	@Title Heapify
 	@Description: to ensure the nature of heap
 	@Author WuChengPei 2022-04-26 15:18:17
 	@Param node:
*/
func (heap *Heap) Heapify(node int) {
	nums := heap.nums
	if node >= heap.length {
		return
	}
	maxNodeIndex := node
	lChild, rChild := (node*2)+1, (node*2)+2
	if heap.kind == "MAX" {
		if lChild < heap.length && nums[maxNodeIndex] < nums[lChild] {
			maxNodeIndex = lChild
		}
		if rChild < heap.length && nums[maxNodeIndex] < nums[rChild] {
			maxNodeIndex = rChild
		}
		if maxNodeIndex != node {
			heap.SwapNode(node, maxNodeIndex)
			heap.Heapify(maxNodeIndex)
		}
	} else {
		if lChild < heap.length && nums[maxNodeIndex] > nums[lChild] {
			maxNodeIndex = lChild
		}
		if rChild < heap.length && nums[maxNodeIndex] > nums[rChild] {
			maxNodeIndex = rChild
		}
		if maxNodeIndex != node {
			heap.SwapNode(node, maxNodeIndex)
			heap.Heapify(maxNodeIndex)
		}
	}

}
