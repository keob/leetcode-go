package solution1104

func reverse(label, row int) int {
	return 1<<(row-1) + 1<<row - 1 - label
}

func pathInZigZagTree(label int) (res []int) {
	row, rowStart := 1, 1

	for rowStart*2 <= label {
		row++
		rowStart *= 2
	}

	if row%2 == 0 {
		label = reverse(label, row)
	}

	for row > 0 {
		if row%2 == 0 {
			res = append(res, reverse(label, row))
		} else {
			res = append(res, label)
		}
		row--
		label >>= 1
	}

	for i, n := 0, len(res); i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}

	return
}
