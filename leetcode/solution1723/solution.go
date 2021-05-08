package solution1723

import "sort"

func minimumTimeRequired(jobs []int, k int) int {
	sort.Ints(jobs)
	low, high := 0, len(jobs)-1
	for low < high {
		jobs[low], jobs[high] = jobs[high], jobs[low]
		low++
		high--
	}

	l, r := jobs[0], 0
	for _, v := range jobs {
		r += v
	}

	for l < r {
		mid := (l + r) >> 1
		if check(jobs, k, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func check(jobs []int, k int, limit int) bool {
	workloads := make([]int, k)
	return backtrack(jobs, workloads, 0, limit)
}

func backtrack(jobs []int, workloads []int, i int, limit int) bool {
	if i >= len(jobs) {
		return true
	}

	cur := jobs[i]

	for j := range workloads {
		if workloads[j]+cur <= limit {
			workloads[j] += cur
			if backtrack(jobs, workloads, i+1, limit) {
				return true
			}
			workloads[j] -= cur
		}
		if workloads[j] == 0 || workloads[j]+cur == limit {
			break
		}
	}
	return false
}
