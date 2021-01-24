package solution674

func findLengthOfLCIS(nums []int) (ans int) {
	start := 0
	for i, v := range nums {
		if i > 0 && v <= nums[i-1] {
			start = i
		}
		ans = max(ans, i-start+1)
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
