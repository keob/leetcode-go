package solution525

func findMaxLength(nums []int) (res int) {
	m := map[int]int{0: -1}
	counter := 0

	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}
		if prev, has := m[counter]; has {
			res = max(res, i-prev)
		} else {
			m[counter] = i
		}
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
