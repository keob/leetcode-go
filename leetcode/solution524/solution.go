package solution524

func findLongestWord(s string, dictionary []string) (res string) {
	for _, v := range dictionary {
		i := 0
		for j := range s {
			if s[j] == v[i] {
				i++
				if i == len(v) {
					if len(v) > len(res) || len(v) == len(res) && v < res {
						res = v
					}
					break
				}
			}
		}
	}
	return
}
