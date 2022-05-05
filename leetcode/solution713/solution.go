package solution713

func numSubarrayProductLessThanK(nums []int, k int) (res int) {
	for i, j, prev := 0, 0, 1; j < len(nums); j++ {
		prev *= nums[j]
		for i <= j && prev >= k {
			prev, i = prev/nums[i], i+1
		}
		res += j - i + 1
	}
	return
}
