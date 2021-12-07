package solution1034

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	originalColor := grid[row][col]
	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	var dfs func(r, c int) bool
	dfs = func(r, c int) bool {
		if r < 0 || c < 0 || r == m || c == n {
			return true
		}
		if grid[r][c] == 0 || grid[r][c] == -1 {
			return false
		}
		if grid[r][c] != originalColor {
			return true
		}
		grid[r][c] = 0
		cur := false
		for _, dir := range dirs {
			if dfs(r+dir[0], c+dir[1]) {
				cur = true
			}
		}
		if cur {
			grid[r][c] = -1
		}
		return false
	}

	dfs(row, col)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case 0:
				grid[i][j] = originalColor
			case -1:
				grid[i][j] = color
			}
		}
	}

	return grid
}
