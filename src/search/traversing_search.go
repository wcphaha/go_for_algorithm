package search

/*
	@Title TraversingSearch
 	@Description: search value by traversing the sequence, O(n)
 	@Author WuChengPei 2022-04-25 10:58:26
 	@Param sequence: input sequence
 	@Param value: searching value
 	@Return found: whether the value was found
 	@Return index: if the value was found, return the first index in sequence; else return -1
*/
func TraversingSearch(sequence []int, value int) (found bool, index int) {
	if len(sequence) == 0 {
		found, index = false, -1
		return
	}
	for ind, val := range sequence {
		if val == value {
			found, index = true, ind
			return
		}
	}
	found, index = false, -1
	return
}
