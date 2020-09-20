package offer20

import "strings"

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}

	s = strings.TrimSpace(s)

	var status [10][5]int = [10][5]int{
		{1, 2, 3, 8, 8},
		{8, 2, 3, 8, 8},
		{8, 8, 4, 8, 8},
		{8, 4, 3, 5, 9},
		{8, 8, 4, 5, 9},
		{6, 8, 7, 8, 8},
		{8, 8, 7, 8, 8},
		{8, 8, 7, 8, 9},
	}

	cur := 0
	idx := 0

	for _, c := range s {
		if c == '-' || c == '+' {
			idx = 0
		} else if c == '.' {
			idx = 1
		} else if c >= '0' && c <= '9' {
			idx = 2
		} else if c == 'e' || c == 'E' {
			idx = 3
		} else {
			return false
		}
		cur = status[cur][idx]
		if cur == 8 {
			break
		} else if cur == 9 {
			break
		}
	}

	if cur != 8 && cur != 9 {
		cur = status[cur][4]
	}

	if cur == 8 {
		return false
	}

	return true
}
