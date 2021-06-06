package solution474

import "strings"

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		zeros := strings.Count(s, "0")
		ones := len(s) - zeros
		for j := m; j >= zeros; j-- {
			for k := n; k >= ones; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zeros][k-ones]+1)
			}
		}
	}

	return dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
