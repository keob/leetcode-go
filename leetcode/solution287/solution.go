package solution287

func findDuplicate(nums []int) int {
	n := len(nums)
	left, right := 1, n-1
	res := -1
	for left <= right {
		mid := (left + right) >> 1
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			left = mid + 1
		} else {
			right = mid - 1
			res = mid
		}
	}
	return res
}
