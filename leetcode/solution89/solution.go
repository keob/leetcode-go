package solution89

func grayCode(n int) []int {
	res := make([]int, 1<<n)
	for i := range res {
		res[i] = i>>1 ^ i
	}
	return res
}
