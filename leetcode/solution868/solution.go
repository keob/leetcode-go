package solution868

func binaryGap(n int) int {
	res := 0
	for i, last := 0, -1; n > 0; i++ {
		if n&1 == 1 {
			if last != -1 {
				res = max(res, i-last)
			}
			last = i
		}
		n >>= 1
	}
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
