package solution977

func sortedSquares(a []int) []int {
	n := len(a)
	res := make([]int, n)
	i, j := 0, n-1

	for pos := n - 1; pos >= 0; pos-- {
		if x, y := a[i]*a[i], a[j]*a[j]; x > y {
			res[pos] = x
			i++
		} else {
			res[pos] = y
			j--
		}
	}

	return res
}
