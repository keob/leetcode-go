package solution633

func judgeSquareSum(c int) bool {
	for base := 2; base*base <= c; base++ {
		if c%base > 0 {
			continue
		}

		exp := 0
		for ; c%base == 0; c /= base {
			exp++
		}

		if base%4 == 3 && exp%2 != 0 {
			return false
		}
	}

	return c%4 != 3
}
