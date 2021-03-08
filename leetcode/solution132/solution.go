package solution132

func minCut(s string) int {
	if len(s) <= 1 {
		return 0
	}

	n := len(s)
	isPali := make([][]bool, len(s))

	for i := range isPali {
		isPali[i] = make([]bool, n)
		for j := range isPali[i] {
			isPali[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			isPali[i][j] = s[i] == s[j] && isPali[i+1][j-1]
		}
	}

	dp := make([]int, n)

	for i := range dp {
		if isPali[0][i] {
			dp[i] = 0
			continue
		}
		dp[i] = i
		for j := 0; j < i; j++ {
			if isPali[j+1][i] {
				if dp[j]+1 < dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	return dp[n-1]
}
