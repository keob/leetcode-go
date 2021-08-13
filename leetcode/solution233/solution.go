package solution233

func countDigitOne(n int) (res int) {
	for i := 1; i <= n; i *= 10 {
		a := n / (i * 10)
		b := n % (i * 10)
		c := min(max(b-i+1, 0), i)
		res += a*i + c
	}

	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
