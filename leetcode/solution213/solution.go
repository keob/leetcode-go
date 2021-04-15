package solution213

func rob(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	return max(helper(nums[:n-1]), helper(nums[1:]))
}

func helper(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])

	for _, v := range nums[2:] {
		first, second = second, max(first+v, second)
	}

	return second
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
