package solution389

func findTheDifference(s string, t string) (diff byte) {
	for i := range s {
		diff ^= s[i] ^ t[i]
	}

	return diff ^ t[len(t)-1]
}
