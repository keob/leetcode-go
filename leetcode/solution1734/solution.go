package solution1734

func decode(encoded []int) []int {
	n := len(encoded)
	sum1, sum2 := 0, 0

	for i := 1; i <= n+1; i++ {
		sum1 ^= i
	}

	for i := 1; i < n; i += 2 {
		sum2 ^= encoded[i]
	}

	res := make([]int, n+1)
	res[0] = sum1 ^ sum2

	for i, v := range encoded {
		res[i+1] = res[i] ^ v
	}

	return res
}
