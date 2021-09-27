package solution639

func numDecodings(s string) int {
	const mod int = 1e9 + 7
	a, b, c := 0, 1, 0
	for i := range s {
		c = b * checkOneDigit(s[i]) % mod
		if i > 0 {
			c = (c + a*checkTwoDigits(s[i-1], s[i])) % mod
		}
		a, b = b, c
	}
	return c
}

func checkOneDigit(c byte) int {
	if c == '*' {
		return 9
	}
	if c == '0' {
		return 0
	}
	return 1
}

func checkTwoDigits(c1 byte, c2 byte) int {
	if c1 == '*' && c2 == '*' {
		return 15
	}
	if c1 == '*' {
		if c2 <= '6' {
			return 2
		}
		return 1
	}
	if c2 == '*' {
		if c1 == '1' {
			return 9
		}
		if c1 == '2' {
			return 6
		}
		return 0
	}
	if c1 != '0' && (c1-'0')*10+(c2-'0') <= 26 {
		return 1
	}
	return 0
}
