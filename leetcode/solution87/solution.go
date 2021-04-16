package solution87

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	length := len(s1)
	dp := make([][][]bool, length)

	for i := 0; i < length; i++ {
		dp[i] = make([][]bool, length)
		for j := 0; j < length; j++ {
			dp[i][j] = make([]bool, length+1)
		}
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if s1[i] == s2[j] {
				dp[i][j][1] = true
			}
		}
	}

	for len := 2; len <= length; len++ {
		for i := 0; i <= length-len; i++ {
			for j := 0; j <= length-len; j++ {
				for k := 1; k <= len-1; k++ {
					if dp[i][j][k] && dp[i+k][j+k][len-k] {
						dp[i][j][len] = true
						break
					}
					if dp[i][j+len-k][k] && dp[i+k][j][len-k] {
						dp[i][j][len] = true
						break
					}
				}
			}
		}
	}

	return dp[0][0][length]
}
