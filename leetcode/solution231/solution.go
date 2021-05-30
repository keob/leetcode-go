package solution231

func isPowerOfTwo(n int) bool {
	const big = 1 << 30
	return n > 0 && big%n == 0
}
