package solution1004

func longestOnes(A []int, K int) int {
	n := len(A)
	l, r := 0, 0

	for r < n {
		if A[r] == 0 {
			K--
		}
		if K < 0 {
			if A[l] == 0 {
				K++
			}
			l++
		}
		r++
	}

	return r - l
}
