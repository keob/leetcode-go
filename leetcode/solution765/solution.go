package solution765

func minSwapsCouples(row []int) (ans int) {
	n := len(row)
	graph := make([][]int, n/2)

	for i := 0; i < n; i += 2 {
		l, r := row[i]/2, row[i+1]/2
		if l != r {
			graph[l] = append(graph[l], r)
			graph[r] = append(graph[r], l)
		}
	}

	vis := make([]bool, n/2)

	for i, vs := range vis {
		if !vs {
			vis[i] = true
			cnt := 0
			q := []int{i}
			for len(q) > 0 {
				cnt++
				v := q[0]
				q = q[1:]
				for _, w := range graph[v] {
					if !vis[w] {
						vis[w] = true
						q = append(q, w)
					}
				}
			}
			ans += cnt - 1
		}
	}

	return
}
