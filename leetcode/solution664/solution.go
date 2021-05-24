package solution664

func strangePrinter(s string) int {
	m := make(map[int]int)
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i > j {
			return 0
		}

		if ans, ok := m[i*100+j]; ok {
			return ans
		}
		ans := dp(i+1, j) + 1
		for k := i + 1; k < j+1; k++ {
			if s[k] == s[i] {
				ans = min(ans, dp(i, k-1)+dp(k+1, j))
			}
		}
		m[i*100+j] = ans
		return ans
	}

	return dp(0, len(s)-1)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
