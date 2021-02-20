package solution697

type entry struct {
	beg   int
	end   int
	count int
}

func findShortestSubArray(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	m := make(map[int]entry)

	for i, n := range nums {
		v, ok := m[n]
		if !ok {
			m[n] = entry{beg: i, end: i, count: 1}
			continue
		}
		v.count++
		v.end = i
		m[n] = v
	}

	res := entry{}

	for _, v := range m {
		if v.count > res.count {
			res = v
		} else if v.count == res.count && (v.end-v.beg) < (res.end-res.beg) {
			res = v
		}
	}

	return res.end - res.beg + 1
}
