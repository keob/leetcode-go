package solution282

func addOperators(num string, target int) (res []string) {
	n := len(num)

	var backtrack func(expr []byte, i, ans, mul int)
	backtrack = func(expr []byte, i, ans, mul int) {
		if i == n {
			if ans == target {
				res = append(res, string(expr))
			}
			return
		}
		signIndex := len(expr)
		if i > 0 {
			expr = append(expr, 0)
		}
		for j, val := i, 0; j < n && (j == i || num[i] != '0'); j++ {
			val = val*10 + int(num[j]-'0')
			expr = append(expr, num[j])
			if i == 0 {
				backtrack(expr, j+1, val, val)
			} else {
				expr[signIndex] = '+'
				backtrack(expr, j+1, ans+val, val)
				expr[signIndex] = '-'
				backtrack(expr, j+1, ans-val, -val)
				expr[signIndex] = '*'
				backtrack(expr, j+1, ans-mul+mul*val, mul*val)
			}
		}
	}

	backtrack(make([]byte, 0, n*2-1), 0, 0, 0)

	return
}
