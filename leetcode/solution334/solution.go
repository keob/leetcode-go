package solution334

func increasingTriplet(nums []int) bool {
	m1, m2 := 1<<31-1, 1<<31-1

	for _, v := range nums {
		if m1 >= v {
			m1 = v
		} else if m2 >= v {
			m2 = v
		} else {
			return true
		}
	}

	return false
}
