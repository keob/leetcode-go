package solution149

func maxPoints(points [][]int) (res int) {
	n := len(points)
	if n <= 2 {
		return n
	}

	for i, p := range points {
		if res >= n-i || res > n/2 {
			break
		}
		cnt := map[int]int{}
		for _, q := range points[i+1:] {
			x, y := p[0]-q[0], p[1]-q[1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 {
					x, y = -x, -y
				}
				g := gcd(abs(x), abs(y))
				x /= g
				y /= g
			}
			cnt[y+x*20001]++
		}
		for _, c := range cnt {
			res = max(res, c+1)
		}
	}

	return
}

func gcd(x, y int) int {
	for x != 0 {
		x, y = y%x, x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
