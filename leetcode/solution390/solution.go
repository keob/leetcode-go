package solution390

func lastRemaining(n int) int {
	start := 1
	k, cnt, step := 0, n, 1

	for cnt > 1 {
		if k%2 == 0 {
			start += step
		} else {
			if cnt%2 == 1 {
				start += step
			}
		}
		k++
		cnt >>= 1
		step <<= 1
	}

	return start
}
