package solution2016

func maximumDifference(nums []int) int {
	res := -1
	for i, min := 1, nums[0]; i < len(nums); i++ {
		if nums[i] > min {
			res = max(res, nums[i]-min)
		} else {
			min = nums[i]
		}
	}
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
