package solution797

func allPathsSourceTarget(graph [][]int) (res [][]int) {
	stack := []int{0}
	var dfs func(int)
	dfs = func(x int) {
		if x == len(graph)-1 {
			res = append(res, append([]int(nil), stack...))
			return
		}
		for _, v := range graph[x] {
			stack = append(stack, v)
			dfs(v)
			stack = stack[:len(stack)-1]
		}
	}
	dfs(0)
	return
}
