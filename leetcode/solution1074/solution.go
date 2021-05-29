package solution1074

func subarraySum(nums []int, k int) (res int) {
	m := map[int]int{0: 1}
	n := len(nums)

	for i, pre := 0, 0; i < n; i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			res += m[pre-k]
		}
		m[pre]++
	}

	return
}

func numSubmatrixSumTarget(matrix [][]int, target int) (res int) {
	for i := range matrix {
		sum := make([]int, len(matrix[0]))
		for _, row := range matrix[i:] {
			for c, v := range row {
				sum[c] += v
			}
			res += subarraySum(sum, target)
		}
	}

	return
}
