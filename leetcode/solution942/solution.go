package solution942

func diStringMatch(s string) []int {
	n := len(s)
	res := make([]int, n+1)
	lo, hi := 0, n
	for i, ch := range s {
		if ch == 'I' {
			res[i] = lo
			lo++
		} else {
			res[i] = hi
			hi--
		}
	}
	res[n] = lo
	return res
}
