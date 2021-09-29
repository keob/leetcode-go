package solution517

func findMinMoves(machines []int) (res int) {
	total := 0

	for _, v := range machines {
		total += v
	}

	n := len(machines)
	if total%n > 0 {
		return -1
	}

	avg := total / n
	sum := 0

	for _, num := range machines {
		num -= avg
		sum += num
		res = max(res, max(abs(sum), num))
	}

	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
