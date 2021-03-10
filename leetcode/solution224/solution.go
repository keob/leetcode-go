package solution224

func calculate(s string) int {
	stack := []int{1}
	res := 0
	num := 0
	op := 1

	for i := range s {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
			continue
		}

		res += op * num
		num = 0

		switch ch {
		case '+':
			op = stack[len(stack)-1]
		case '-':
			op = -stack[len(stack)-1]
		case '(':
			stack = append(stack, op)
		case ')':
			stack = stack[:len(stack)-1]
		}
	}

	res += op * num

	return res
}
