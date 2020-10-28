package solution1207

func uniqueOccurrences(arr []int) bool {
	m1 := make([]int, 2002)
	for _, v := range arr {
		m1[v+1000]++
	}

	m2 := make([]bool, 1001)
	for _, v := range m1 {
		if m2[v] {
			return false
		}
		if v != 0 {
			m2[v] = true
		}
	}

	return true
}
