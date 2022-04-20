package solution388

func lengthLongestPath(input string) (res int) {
	stack := []int{}

	for i, n := 0, len(input); i < n; {
		depth := 1
		for ; i < n && input[i] == '\t'; i++ {
			depth++
		}

		length, isFile := 0, false
		for ; i < n && input[i] != '\n'; i++ {
			if input[i] == '.' {
				isFile = true
			}
			length++
		}
		i++

		for len(stack) >= depth {
			stack = stack[:len(stack)-1]
		}

		if (len(stack)) > 0 {
			length += stack[len(stack)-1] + 1
		}
		if isFile {
			res = max(res, length)
		} else {
			stack = append(stack, length)
		}
	}

	return
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
