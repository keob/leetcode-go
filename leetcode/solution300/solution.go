package solution300

func lengthOfLIS(nums []int) int {
	res := []int{}

	for _, num := range nums {
		if len(res) == 0 || res[len(res)-1] < num {
			res = append(res, num)
		} else {
			l, r := 0, len(res)-1
			pos := r
			for l <= r {
				mid := (l + r) >> 1
				if res[mid] >= num {
					pos = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			res[pos] = num
		}
	}

	return len(res)
}
