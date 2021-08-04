package solution611

import "sort"

func triangleNumber(nums []int) (res int) {
	n := len(nums)
	sort.Ints(nums)

	for i, v := range nums {
		k := i
		for j := i + 1; j < n; j++ {
			for k+1 < n && nums[k+1] < v+nums[j] {
				k++
			}
			res += max(k-j, 0)
		}
	}

	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
