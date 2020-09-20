package solution6

func convert(s string, numRows int) string {
	matrix := make([][]rune, numRows)

	for i := 0; i < numRows; i++ {
		matrix[i] = make([]rune, len(s))
	}

	row := 0
	col := 0
	stat := 0

	for _, ch := range s {
		matrix[row][col] = ch
		switch stat {
		case 0:
			if row >= numRows-1 {
				if row <= 0 {
					col++
					continue
				}
				stat = 1
				row--
				col++
			} else {
				row++
			}
		case 1:
			if row == 0 {
				stat = 0
				row++
			} else {
				row--
				col++
			}
		}
	}

	var ret string

	for y := 0; y < numRows; y++ {
		for x := 0; x < len(s); x++ {
			if matrix[y][x] != 0 {
				ret += string(matrix[y][x])
			}
		}
	}

	return ret
}
