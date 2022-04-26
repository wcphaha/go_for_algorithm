package sort

/*
	@Title InsertionSort
 	@Description: implementation of insertion sort
 	@Author WuChengPei 2022-04-25 19:31:05
 	@Param nums:
 	@Return []int:
*/
func InsertionSort(nums []int) []int {
	for i := range nums {
		preIndex := i - 1
		currentValue := nums[i]
		for preIndex >= 0 && nums[preIndex] > currentValue {
			nums[preIndex+1] = nums[preIndex]
			preIndex--
		}
		nums[preIndex+1] = currentValue
	}
	return nums
}
