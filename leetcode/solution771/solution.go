package solution771

func numJewelsInStones(J string, S string) int {
	count := 0
	m := make(map[rune]int)

	for _, v := range J {
		m[v] = 1
	}

	for _, v := range S {
		if _, ok := m[v]; ok {
			count++
		}
	}

	return count
}
