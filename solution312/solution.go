package solution312

func maxCoins(nums []int) int {
	n := len(nums)
	res := make([][]int, n+2)
	val := make([]int, n+2)

	for i := 0; i < n+2; i++ {
		res[i] = make([]int, n+2)
	}

	val[0], val[n+1] = 1, 1

	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += res[i][k] + res[k][j]
				res[i][j] = max(res[i][j], sum)
			}
		}
	}
	return res[0][n+1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
