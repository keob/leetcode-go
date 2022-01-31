package solution1342

func numberOfSteps(num int) (res int) {
	for ; num > 0; num >>= 1 {
		res += num & 1
		if num > 1 {
			res++
		}
	}
	return
}
