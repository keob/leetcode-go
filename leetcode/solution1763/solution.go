package solution1763

import "unicode"

func longestNiceSubstring(s string) (res string) {
	if s == "" {
		return
	}

	lower, upper := 0, 0

	for _, ch := range s {
		if unicode.IsLower(ch) {
			lower |= 1 << (ch - 'a')
		} else {
			upper |= 1 << (ch - 'A')
		}
	}

	if lower == upper {
		return s
	}

	valid := lower & upper

	for i := 0; i < len(s); i++ {
		start := i
		for i < len(s) && valid>>(unicode.ToLower(rune(s[i]))-'a')&1 == 1 {
			i++
		}
		if t := longestNiceSubstring(s[start:i]); len(t) > len(res) {
			res = t
		}
	}

	return
}
