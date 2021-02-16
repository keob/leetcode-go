package solution561

import "sort"

func arrayPairSum(nums []int) (ans int) {
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n; i += 2 {
		ans += nums[i]
	}

	return
}
