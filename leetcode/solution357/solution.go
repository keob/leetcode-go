package solution357

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	res, cur := 10, 9
	for i := 0; i < n-1; i++ {
		cur *= 9 - i
		res += cur
	}
	return res
}
