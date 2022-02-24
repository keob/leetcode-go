package solution1706

func findBall(grid [][]int) []int {
	n := len(grid[0])
	res := make([]int, n)

	for j := range res {
		col := j
		for _, row := range grid {
			dir := row[col]
			col += dir
			if col < 0 || col == n || row[col] != dir {
				col = -1
				break
			}
		}
		res[j] = col
	}

	return res
}
