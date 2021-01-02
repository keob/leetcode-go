package solution239

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	prefixMax := make([]int, n)
	suffixMax := make([]int, n)
	for i, v := range nums {
		if i%k == 0 {
			prefixMax[i] = v
		} else {
			prefixMax[i] = max(prefixMax[i-1], v)
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 || (i+1)%k == 0 {
			suffixMax[i] = nums[i]
		} else {
			suffixMax[i] = max(suffixMax[i+1], nums[i])
		}
	}

	ans := make([]int, n-k+1)
	for i := range ans {
		ans[i] = max(suffixMax[i], prefixMax[i+k-1])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
