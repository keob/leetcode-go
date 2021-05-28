package solution477

func totalHammingDistance(nums []int) (res int) {
	n := len(nums)

	for i := 0; i < 30; i++ {
		c := 0
		for _, val := range nums {
			c += val >> i & 1
		}
		res += c * (n - c)
	}

	return
}
