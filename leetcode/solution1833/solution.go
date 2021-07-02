package solution1833

func maxIceCream(costs []int, coins int) (res int) {
	const mx int = 1e5
	freq := [mx + 1]int{}

	for _, c := range costs {
		freq[c]++
	}

	for i := 1; i <= mx && coins >= i; i++ {
		c := min(freq[i], coins/i)
		res += c
		coins -= i * c
	}

	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
