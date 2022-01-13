package solution747

func dominantIndex(nums []int) int {
	m1, m2, index := -1, -1, 0
	for i, num := range nums {
		if num > m1 {
			m1, m2, index = num, m1, i
		} else if num > m2 {
			m2 = num
		}
	}
	if m1 >= m2*2 {
		return index
	}
	return -1
}
