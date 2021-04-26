package solution1011

func shipWithinDays(weights []int, D int) int {
	sum, max := 0, -1

	for _, weight := range weights {
		sum += weight
		if weight > max {
			max = weight
		}
	}

	l, r, min := max, sum, sum

	for l <= r {
		mid := (r-l)/2 + l
		days := countDays(mid, weights)
		if days <= D {
			if mid < min {
				min = mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return min
}

func countDays(ship int, weights []int) int {
	cur, days := ship, 1

	for _, weight := range weights {
		if cur >= weight {
			cur -= weight
		} else {
			cur = ship - weight
			days++
		}
	}

	return days
}
