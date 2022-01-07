package solution1614

func maxDepth(s string) (res int) {
	size := 0

	for _, ch := range s {
		switch ch {
		case '(':
			size++
			if size > res {
				res = size
			}
		case ')':
			size--
		}
	}

	return res
}
