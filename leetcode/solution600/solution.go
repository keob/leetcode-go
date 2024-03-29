package solution600

func findIntegers(n int) (res int) {
	dp := [31]int{1, 1}

	for i := 2; i < 31; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	for i, pre := 29, 0; i >= 0; i-- {
		val := 1 << i
		if n&val > 0 {
			res += dp[i+1]
			if pre == 1 {
				break
			}
			pre = 1
		} else {
			pre = 0
		}
		if i == 0 {
			res++
		}
	}
	return
}
