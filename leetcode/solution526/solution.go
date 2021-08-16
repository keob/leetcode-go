package solution526

func countArrangement(n int) (res int) {
	vis := make([]bool, n+1)
	match := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i%j == 0 || j%i == 0 {
				match[i] = append(match[i], j)
			}
		}
	}

	var backstack func(int)
	backstack = func(index int) {
		if index > n {
			res++
			return
		}
		for _, x := range match[index] {
			if !vis[x] {
				vis[x] = true
				backstack(index + 1)
				vis[x] = false
			}
		}
	}

	backstack(1)

	return
}
