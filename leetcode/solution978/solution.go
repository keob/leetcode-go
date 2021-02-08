package solution978

func maxTurbulenceSize(arr []int) int {
	ans := 1
	n := len(arr)
	dp0, dp1 := 1, 1

	for i := 1; i < n; i++ {
		if arr[i-1] > arr[i] {
			dp0, dp1 = dp1+1, 1
		} else if arr[i-1] < arr[i] {
			dp0, dp1 = 1, dp0+1
		} else {
			dp0, dp1 = 1, 1
		}
		ans = max(ans, max(dp0, dp1))
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
