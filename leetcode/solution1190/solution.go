package solution1190

func reverseParentheses(s string) string {
	n := len(s)
	pair := make([]int, n)
	stack := []int{}
	res := []byte{}

	for i, b := range s {
		if b == '(' {
			stack = append(stack, i)
		} else if b == ')' {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pair[i], pair[j] = j, i
		}
	}

	for i, step := 0, 1; i < n; i += step {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			step = -step
		} else {
			res = append(res, s[i])
		}
	}

	return string(res)
}
