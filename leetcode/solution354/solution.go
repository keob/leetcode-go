package solution354

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	res := []int{}

	for _, e := range envelopes {
		m := e[1]
		if i := sort.SearchInts(res, m); i < len(res) {
			res[i] = m
		} else {
			res = append(res, m)
		}
	}

	return len(res)
}
