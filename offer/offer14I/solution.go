package offer14I

func cuttingRope(n int) int {
	dp := make(map[int]int)
	dp[1] = 1

	for i := 2; i < n+1; i++ {
		for j := 1; j < (i/2 + 1); j++ {
			dp[i] = max(dp[i], max(j, dp[j])*max(i-j, dp[i-j]))
		}
	}

	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
