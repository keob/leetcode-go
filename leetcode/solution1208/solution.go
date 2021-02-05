package solution1208

func equalSubstring(s string, t string, maxCost int) (res int) {
	n := len(s)
	diff := make([]int, n)

	for i, ch := range s {
		diff[i] = abs(int(ch) - int(t[i]))
	}

	sum, start := 0, 0

	for end, d := range diff {
		sum += d
		for sum > maxCost {
			sum -= diff[start]
			start++
		}
		res = max(res, end-start+1)
	}

	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
