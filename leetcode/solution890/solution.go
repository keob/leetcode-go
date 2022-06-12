package solution890

func findAndReplacePattern(words []string, pattern string) (res []string) {
	for _, word := range words {
		if match(word, pattern) && match(pattern, word) {
			res = append(res, word)
		}
	}
	return
}

func match(word, pattern string) bool {
	mp := map[rune]byte{}
	for i, x := range word {
		y := pattern[i]
		if mp[x] == 0 {
			mp[x] = y
		} else if mp[x] != y {
			return false
		}
	}
	return true
}
