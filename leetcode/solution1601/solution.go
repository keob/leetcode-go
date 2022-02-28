package solution1601

func maximumRequests(n int, requests [][]int) (res int) {
	delta := make([]int, n)
	cnt, zero := 0, n

	var dfs func(int)
	dfs = func(pos int) {
		if pos == len(requests) {
			if zero == n && cnt > res {
				res = cnt
			}
			return
		}

		dfs(pos + 1)

		z := zero
		cnt++
		r := requests[pos]
		x, y := r[0], r[1]
		if delta[x] == 0 {
			zero--
		}
		delta[x]--
		if delta[x] == 0 {
			zero++
		}
		if delta[y] == 0 {
			zero--
		}
		delta[y]++
		if delta[y] == 0 {
			zero++
		}
		dfs(pos + 1)
		delta[y]--
		delta[x]++
		cnt--
		zero = z
	}

	dfs(0)

	return
}
