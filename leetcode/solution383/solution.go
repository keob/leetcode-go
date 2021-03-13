package solution383

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[int32]int)

	for _, v := range magazine {
		m[v] += 1
	}

	for _, v := range ransomNote {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}

	return true
}
