package solution883

func projectionArea(grid [][]int) (res int) {
	for i := 0; i < len(grid); i++ {
		maxRow, maxCol := 0, 0
		for j := 0; j < len(grid); j++ {
			if grid[i][j] != 0 {
				res++
			}
			maxRow = max(maxRow, grid[i][j])
			maxCol = max(maxCol, grid[j][i])
		}
		res += maxRow + maxCol
	}
	return
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
