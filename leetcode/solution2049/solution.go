package solution2049

func countHighestScoreNodes(parents []int) (res int) {
	n := len(parents)
	children := make([][]int, n)
	for node := 1; node < n; node++ {
		p := parents[node]
		children[p] = append(children[p], node)
	}

	maxScore := 0
	var dfs func(int) int
	dfs = func(node int) int {
		score, size := 1, n-1
		for _, ch := range children[node] {
			sz := dfs(ch)
			score *= sz
			size -= sz
		}
		if node > 0 {
			score *= size
		}
		if score == maxScore {
			res++
		} else if score > maxScore {
			maxScore = score
			res = 1
		}
		return n - size
	}
	dfs(0)
	return
}
