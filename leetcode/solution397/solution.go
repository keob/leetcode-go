package solution397

func integerReplacement(n int) (res int) {
	for n != 1 {
		switch {
		case n%2 == 0:
			res++
			n /= 2
		case n%4 == 1:
			res += 2
			n /= 2
		case n == 3:
			res += 2
			n = 1
		default:
			res += 2
			n = n/2 + 1
		}
	}
	return
}
