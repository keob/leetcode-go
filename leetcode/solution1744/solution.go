package solution1744

func canEat(candiesCount []int, queries [][]int) []bool {
	m, n := len(candiesCount), len(queries)
	res := make([]bool, n)
	prefix := make([]int, m+1)

	for i := 1; i <= m; i++ {
		prefix[i] += prefix[i-1] + candiesCount[i-1]
	}

	for i := 0; i < n; i++ {
		ft, fd, dc := queries[i][0], queries[i][1], queries[i][2]
		if (fd+1)*dc <= prefix[ft] {
			res[i] = false
		} else if (fd + 1) <= prefix[ft+1] {
			res[i] = true
		}
	}

	return res
}
