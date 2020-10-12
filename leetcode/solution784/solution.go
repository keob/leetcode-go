package solution784

var res []string

func letterCasePermutation(S string) []string {
	res = []string{}
	arr := make([]rune, 0)

	for _, v := range S {
		arr = append(arr, v)
	}

	dfs(arr, 0, make([]rune, 0))

	return res
}

func dfs(str []rune, depth int, path []rune) {
	if len(str) == depth {
		n1 := make([]rune, depth)
		copy(n1, path)
		res = append(res, string(n1))
		return
	}

	if str[depth] >= 48 && str[depth] <= 57 {
		path = append(path, str[depth])
		dfs(str, depth+1, path)
	} else if str[depth] >= 65 && str[depth] <= 122 {
		path = append(path, str[depth])
		dfs(str, depth+1, path)
		path = path[:len(path)-1]

		path = append(path, str[depth]^32)
		dfs(str, depth+1, path)
	}
}
