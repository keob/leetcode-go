package solution1838

import "sort"

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	res := 1

	for l, r, total := 0, 1, 0; r < len(nums); r++ {
		total += (nums[r] - nums[r-1]) * (r - l)
		for total > k {
			total -= nums[r] - nums[l]
			l++
		}
		res = max(res, r-l+1)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
