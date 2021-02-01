package solution888

import (
	"sort"
)

func fairCandySwap(A []int, B []int) []int {
	sumA := 0
	sumB := 0
	for _, v := range A {
		sumA += v
	}

	for _, v := range B {
		sumB += v
	}

	sort.Ints(A)
	sort.Ints(B)

	target := (sumA - sumB) / 2
	i := 0
	j := 0

	for i < len(A) && j < len(B) {

		if A[i] == B[j]+target {
			return []int{A[i], B[j]}
		}

		if A[i] < B[j]+target {
			i++
		} else {
			j++
		}

	}

	return []int{0, 0}
}
