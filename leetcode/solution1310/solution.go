package solution1310

func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, len(arr)+1)
	ans := make([]int, len(queries))

	for i, v := range arr {
		xors[i+1] = xors[i] ^ v
	}

	for i, v := range queries {
		ans[i] = xors[v[0]] ^ xors[v[1]+1]
	}

	return ans
}
