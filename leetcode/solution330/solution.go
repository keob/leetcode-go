package solution330

func minPatches(nums []int, n int) (res int) {
	for i, j := 0, 1; j <= n; {
		if i < len(nums) && nums[i] <= j {
			j += nums[i]
			i++
		} else {
			j *= 2
			res++
		}
	}
	return
}
