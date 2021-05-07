package solution1486

func xorOperation(n int, start int) int {
	s, e := start>>1, n&start&1
	res := sumXor(s-1) ^ sumXor(s+n-1)

	return res<<1 | e
}

func sumXor(x int) int {
	switch x % 4 {
	case 0:
		return x
	case 1:
		return 1
	case 2:
		return x + 1
	default:
		return 0
	}
}
