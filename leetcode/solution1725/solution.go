package solution1725

func countGoodRectangles(rectangles [][]int) (res int) {
	maxLen := 0
	for _, rect := range rectangles {
		k := min(rect[0], rect[1])
		if k == maxLen {
			res++
		} else if k > maxLen {
			maxLen, res = k, 1
		}
	}
	return
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
