package solution454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	result := 0
	sums := make(map[int]int)

	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			sums[C[i]+D[j]]++
		}
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			result += sums[-A[i]-B[j]]
		}
	}

	return result
}
