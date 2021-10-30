package solution260

func singleNumber(nums []int) []int {
	xorSum := 0

	for _, num := range nums {
		xorSum ^= num
	}

	l := xorSum & -xorSum
	a, b := 0, 0

	for _, num := range nums {
		if num&l > 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}
