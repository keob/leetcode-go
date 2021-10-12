package solution29

func divide(dividend int, divisor int) (res int) {
	sign := (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0)
	dividend = abs(dividend)
	divisor = abs(divisor)

	for dividend >= divisor {
		d, q := divisor, 1
		for d <= dividend {
			d, q = d<<1, q<<1
		}
		res += q >> 1
		dividend = dividend - (d >> 1)
	}

	if sign {
		res = ^res + 1
	}

	if res < -2147483648 || res > 2147483647 {
		return 2147483647
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
