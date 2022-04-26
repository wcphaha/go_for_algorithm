package sort

type Heap struct {
	kind   string
	nums   []int
	length int
}

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
