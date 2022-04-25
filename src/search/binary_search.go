package search

/*
	@Title BinarySearch
 	@Description: implementation of binary-search
 	@Author WuChengPei 2022-04-24 21:19:01
 	@Param sequence: input sequence, must be orderly(here ascending order)
 	@Param value: searching value
 	@Return found: whether the value was found
*/
func BinarySearch(sequence []int, value int) (found bool) {
	var mid, left, right int
	left, right = 0, len(sequence)-1
	for left <= right {
		mid = (left + right) / 2
		if sequence[mid] == value {
			found = true
			return
		}
		if value > sequence[mid] {
			left = mid + 1
			continue
		}
		if value < sequence[mid] {
			right = mid - 1
			continue
		}
	}
	found = false
	return
}
