package solution778

func swimInWater(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	var a = []int{-1, 1, 0, 0}
	var b = []int{0, 0, -1, 1}
	result := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		result[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			result[i][j] = 50 * 50
		}
	}

	var queue [][]int
	queue = append(queue, []int{0, 0})
	result[0][0] = grid[0][0]
	n := len(grid)

	for len(queue) > 0 {
		row := queue[0][0]
		col := queue[0][1]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			newRow := row + a[i]
			newCol := col + b[i]
			if newRow >= 0 && newRow < n && newCol >= 0 && newCol < n {
				maxNum := max(result[row][col], grid[newRow][newCol])
				if maxNum < result[newRow][newCol] {
					result[newRow][newCol] = maxNum
					queue = append(queue, []int{newRow, newCol})
				}
			}
		}
	}

	return result[n-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
