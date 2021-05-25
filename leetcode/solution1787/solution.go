package solution1787

func minChanges(nums []int, k int) int {
	const scope = 1 << 10
	const inf = 1<<31 - 1

	n := len(nums)
	f := make([]int, scope)

	for i := range f {
		f[i] = inf
	}

	f[0] = 0

	for i := 0; i < k; i++ {
		cnt := make(map[int]int)
		size := 0
		for j := i; j < n; j += k {
			cnt[nums[j]]++
			size++
		}

		minT := min(f...)

		g := make([]int, scope)

		for j := range g {
			g[j] = minT
		}

		for mask := range g {
			for x, cntX := range cnt {
				g[mask] = min(g[mask], f[mask^x]-cntX)
			}
		}

		for j := range g {
			g[j] += size
		}

		f = g
	}

	return f[0]
}

func min(arr ...int) int {
	res := arr[0]
	for _, v := range arr[1:] {
		if v < res {
			res = v
		}
	}

	return res
}
