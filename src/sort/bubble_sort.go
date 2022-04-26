package sort

/*
	@Title BubbleSort
 	@Description: implementation of bubble-sort
 	@Author WuChengPei 2022-04-26 15:30:54
 	@Param nums:
 	@Return []int:
*/
func BubbleSort(nums []int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
