package solution58

func lengthOfLastWord(s string) (res int) {
	index := len(s) - 1

	for s[index] == ' ' {
		index--
	}

	for index >= 0 && s[index] != ' ' {
		res++
		index--
	}
	return
}
