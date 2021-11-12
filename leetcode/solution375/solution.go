package solution375

func getMoneyAmount(n int) int {
	f := make([][]int, n+1)

	for i := range f {
		f[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			minCost := 1<<31 - 1
			for k := i; k < j; k++ {
				cost := k + max(f[i][k-1], f[k+1][j])
				if cost < minCost {
					minCost = cost
				}
			}
			f[i][j] = minCost
		}
	}

	return f[1][n]
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
