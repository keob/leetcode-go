package solution1047

func removeDuplicates(S string) string {
	stack := []byte{}

	for i := range S {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}

	return string(stack)
}
