package solution228

import "strconv"

func summaryRanges(nums []int) []string {
	ans := make([]string, 0)
	n := len(nums)

	if n == 0 {
		return ans
	}

	start, end := nums[0], nums[0]

	for i, j := range nums {
		if i == n-1 || j != nums[i+1]-1 {
			end = nums[i]
			if start == end {
				ans = append(ans, strconv.Itoa(end))
			} else {
				ans = append(ans, strconv.Itoa(start)+"->"+strconv.Itoa(end))
			}
			if i != n-1 {
				start = nums[i+1]
			}
		}
	}

	return ans
}
