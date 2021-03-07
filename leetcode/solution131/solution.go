package solution131

func partition(s string) (ans [][]string) {
	n := len(s)
	dp := make([][]bool, n)

	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}

	splits := []string{}

	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if dp[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}

	dfs(0)

	return
}
