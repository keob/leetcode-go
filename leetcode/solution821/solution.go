package solution821

func shortestToChar(s string, c byte) []int {
	n := len(s)
	res := make([]int, n)
	for i, last := 0, -n; i < n; i++ {
		if s[i] == c {
			if last == -n {
				for j := i; j >= 0; j-- {
					res[j] = i - j
				}
			} else {
				for j := i; j > (i+last-1)>>1; j-- {
					res[j] = i - j
				}
			}
			last = i
		} else {
			res[i] = i - last
		}
	}
	return res
}
