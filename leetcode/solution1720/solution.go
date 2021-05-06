package solution1720

func decode(encoded []int, first int) []int {
	ans := make([]int, len(encoded)+1)
	ans[0] = first

	for k, v := range encoded {
		ans[k+1] = ans[k] ^ v
	}

	return ans
}
