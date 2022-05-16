package solution1021

import "strings"

func removeOuterParentheses(s string) string {
	var b strings.Builder
	var cnt int

	for _, c := range s {
		if c == '(' {
			if cnt > 0 {
				b.WriteRune(c)
			}
			cnt++
		} else {
			if cnt > 1 {
				b.WriteRune(c)
			}
			cnt--
		}
	}
	return b.String()
}
