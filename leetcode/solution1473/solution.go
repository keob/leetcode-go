package solution1473

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	dp := make([][][]int, m+1)
	res := 1<<31 - 1

	for i := range dp {
		dp[i] = make([][]int, n+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, target+1)
			dp[i][j][0] = 1<<31 - 1
		}
	}

	for i := 1; i <= m; i++ {
		color := houses[i-1]
		for j := 1; j <= n; j++ {
			for k := 1; k <= target; k++ {
				if k > i {
					dp[i][j][k] = 1<<31-1
					continue
				}
				if color == 0 {
					c := cost[i-1][j-1]
					temp1 := 1<<31-1
					for p := 1; p <= n; p++ {
						if p != j {
							temp1 = min(temp1, dp[i-1][p][k-1])
						}
					}
					temp2 := dp[i-1][j][k]
					dp[i][j][k] = min(temp1, temp2) + c
				} else {
					if color == j {
						temp1 := 1<<31-1
						for p := 1; p <= n; p++ {
							if p != j {
								temp1 = min(temp1, dp[i-1][p][k-1])
							}
						}
						temp2 := dp[i-1][j][k]
						dp[i][j][k] = min(temp1, temp2)
					} else {
						dp[i][j][k] = 1<<31 - 1
					}
				}
			}
		}
	}

	for i := 1; i <= n; i++ {
		res = min(res, dp[m][i][target])
	}

	if res == 1<<31-1 {
		return -1
	}

	return res
}

func min(args ...int) int {
	min := args[0]

	for _, i := range args {
		if i < min {
			min = i
		}
	}

	return min
}
