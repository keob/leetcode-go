package solution290

import "strings"

func wordPattern(pattern string, str string) bool {
	strArr := strings.Fields(str)

	if len(pattern) != len(strArr) {
		return false
	}

	hash1 := make(map[byte]string)
	hash2 := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		v, ok := hash1[pattern[i]]
		v2, ok2 := hash2[strArr[i]]
		if ok && v != strArr[i] || ok2 && v2 != pattern[i] {
			return false
		} else {
			hash1[pattern[i]] = strArr[i]
			hash2[strArr[i]] = pattern[i]
		}
	}

	return true
}
