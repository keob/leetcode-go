package solution275

func hIndex(citations []int) int {
	n := len(citations)
	if n == 0 {
		return 0
	}

	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		fi := citations[mid] + mid - n
		if fi >= 0 {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return n - l
}
