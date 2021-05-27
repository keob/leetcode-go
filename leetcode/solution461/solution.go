package solution461

func hammingDistance(x int, y int) (res int) {
	for s := x ^ y; s > 0; s >>= 1 {
		res += s & 1
	}

	return
}
