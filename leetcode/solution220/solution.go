package solution220

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	m := map[int]int{}
	for i, x := range nums {
		id := key(x, t+1)
		if _, has := m[id]; has {
			return true
		}
		if y, has := m[id-1]; has && abs(x-y) <= t {
			return true
		}
		if y, has := m[id+1]; has && abs(x-y) <= t {
			return true
		}
		m[id] = x
		if i >= k {
			delete(m, key(nums[i-k], t+1))
		}
	}
	return false
}

func key(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
