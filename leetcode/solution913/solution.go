package solution913

func catMouseGame(graph [][]int) int {
	n := len(graph)
	dp := make([][][2]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, n)
		for j := 0; j < n; j++ {
			dp[i][j][0] = -1
			dp[i][j][1] = 100
		}
	}

	var dfs func(i, j, times int) int
	dfs = func(i, j, times int) int {
		if dp[i][j][0] != -1 && dp[i][j][1] <= times {
			return dp[i][j][0]
		}
		if i == 0 {
			return 1
		}
		if i == j {
			return 2
		}
		dp[i][j][0] = 0
		dp[i][j][1] = times
		ans := 2
		for _, ratPos := range graph[i] {
			result := 1
			if ratPos == j {
				result = 2
				continue
			}
			for _, catPos := range graph[j] {
				if catPos == 0 {
					continue
				}
				res := dfs(ratPos, catPos, times+1)
				if res == 2 {
					result = 2
					break
				} else if res == 0 && result != 2 {
					result = 0
				}
			}
			if result == 1 {
				ans = 1
				break
			} else if result == 0 && ans != 1 {
				ans = 0
			}
		}
		dp[i][j][0] = ans
		return ans
	}

	return dfs(1, 2, 0)
}
