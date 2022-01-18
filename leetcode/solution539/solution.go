package solution539

import "sort"

func findMinDifference(timePoints []string) int {
	if len(timePoints) > 1440 {
		return 0
	}

	sort.Strings(timePoints)
	res := 1<<31 - 1
	t0Minutes := getMinutes(timePoints[0])
	preMinutes := t0Minutes

	for _, t := range timePoints[1:] {
		minutes := getMinutes(t)
		res = min(res, minutes-preMinutes)
		preMinutes = minutes
	}

	res = min(res, t0Minutes+1440-preMinutes)

	return res
}

func getMinutes(time string) int {
	return (int(time[0]-'0')*10+int(time[1]-'0'))*60 +
		int(time[3]-'0')*10 + int(time[4]-'0')
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
