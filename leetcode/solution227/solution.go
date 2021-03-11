package solution227

func calculate(s string) (res int) {
	stack := []int{}
	preOp := '+'
	num := 0

	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preOp {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preOp = ch
			num = 0
		}
	}

	for _, v := range stack {
		res += v
	}

	return
}
