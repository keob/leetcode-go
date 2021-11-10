package solution495

func findPoisonedDuration(timeSeries []int, duration int) int {
	res := duration

	for i := 0; i < len(timeSeries)-1; i++ {
		if timeSeries[i+1] >= timeSeries[i]+duration {
			res += duration
		} else {
			res += timeSeries[i+1] - timeSeries[i]
		}
	}

	return res
}
