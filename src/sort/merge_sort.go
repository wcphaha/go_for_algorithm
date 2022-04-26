package sort

/*
	@Title MergeSort
 	@Description: merge sort
 	@Author WuChengPei 2022-04-25 20:00:23
 	@Param nums:
 	@Return []int:
*/
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	nums1 := nums[:mid]
	nums2 := nums[mid:]
	return merge(MergeSort(nums1), MergeSort(nums2))
}
func merge(nums1, nums2 []int) []int {
	p, q := 0, 0
	var res []int
	for p < len(nums1) && q < len(nums2) {
		if nums1[p] < nums2[q] {
			res = append(res, nums1[p])
			p++
		} else {
			res = append(res, nums2[q])
			q++
		}
	}
	if p == len(nums1) && q == len(nums2) {
		return res
	} else {
		if p == len(nums1) {
			res = append(res, nums2[q:]...)
		}
		if q == len(nums2) {
			res = append(res, nums1[p:]...)
		}
	}
	return res
}
