package solution728

func selfDividingNumbers(left int, right int) (res []int) {
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			res = append(res, i)
		}
	}
	return res
}

func isSelfDividing(a int) bool {
	if a < 10 {
		return true
	}

	n := a
	for n > 0 {
		r := n % 10
		if r == 0 || a%r != 0 {
			return false
		}
		n /= 10
	}

	return true
}
