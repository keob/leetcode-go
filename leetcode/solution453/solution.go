package solution453

func minMoves(nums []int) (res int) {
	min := nums[0]

	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}

	for _, num := range nums {
		res += num - min
	}

	return
}
