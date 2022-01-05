package solution1576

func modifyString(s string) string {
	res := make([]rune, len(s))

	for i, r := range s {
		if s[i] == byte('?') {
			for j := 0; j < 26; j++ {
				b := rune('a' + j)
				if i > 0 && res[i-1] == b {
					continue
				}
				if i < len(s)-1 && rune(s[i+1]) == b {
					continue
				}
				res[i] = b
				break
			}
		} else {
			res[i] = r
		}
	}

	return string(res)
}
