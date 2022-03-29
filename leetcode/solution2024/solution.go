package solution2024

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutiveChar(answerKey, k, 'T'),
		maxConsecutiveChar(answerKey, k, 'F'))
}

func maxConsecutiveChar(answerKey string, k int, ch byte) int {
	left, count, res := 0, 0, 0
	for right := range answerKey {
		if answerKey[right] != ch {
			count++
		}
		for count > k {
			if answerKey[left] != ch {
				count--
			}
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
