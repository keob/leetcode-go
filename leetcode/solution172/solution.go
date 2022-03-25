package solution172

func trailingZeroes(n int) (res int) {
	for n > 0 {
		n /= 5
		res += n
	}
	return
}
