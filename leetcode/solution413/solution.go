package solution413

func numberOfArithmeticSlices(nums []int) (res int) {
	n := len(nums)
	if n == 1 {
		return
	}

	d, t := nums[0]-nums[1], 0

	for i := 2; i < n; i++ {
		if nums[i-1]-nums[i] == d {
			t++
		} else {
			d, t = nums[i-1]-nums[i], 0
		}
		res += t
	}

	return
}
