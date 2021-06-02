package solution523

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	m := map[int]int{0: -1}
	remainder := 0

	for i, num := range nums {
		remainder = (remainder + num) % k
		if prev, has := m[remainder]; has {
			if i-prev >= 2 {
				return true
			}
		} else {
			m[remainder] = i
		}
	}
	return false
}
