package solution482

import "unicode"

func licenseKeyFormatting(s string, k int) string {
	res := []byte{}

	for i, cnt := len(s)-1, 0; i >= 0; i-- {
		if s[i] != '-' {
			res = append(res, byte(unicode.ToUpper(rune(s[i]))))
			cnt++
			if cnt%k == 0 {
				res = append(res, '-')
			}
		}
	}

	if len(res) > 0 && res[len(res)-1] == '-' {
		res = res[:len(res)-1]
	}

	for i, n := 0, len(res); i < n/2; i++ {
		res[i], res[n-i-1] = res[n-1-i], res[i]
	}

	return string(res)
}
