package solution995

func minKBitFlips(A []int, K int) (ans int) {
	n := len(A)
	cnt := 0

	for i, v := range A {
		if i >= K && A[i-K] > 1 {
			cnt ^= 1
			A[i-K] -= 2
		}

		if v == cnt {
			if i+K > n {
				return -1
			}
			ans++
			cnt ^= 1
			A[i] += 2
		}
	}

	return
}
