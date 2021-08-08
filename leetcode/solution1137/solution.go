package solution1137

type matrix [3][3]int

func (m matrix) mul(n matrix) matrix {
	res := matrix{}
	for i, row := range m {
		for j := range n[0] {
			for k, v := range row {
				res[i][j] += v * n[k][j]
			}
		}
	}
	return res
}

func (m matrix) pow(n int) matrix {
	res := matrix{}
	for i := range res {
		res[i][i] = 1
	}
	for ; n > 0; n >>= 1 {
		if n&1 > 0 {
			res = res.mul(m)
		}
		m = m.mul(m)
	}
	return res
}

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	m := matrix{
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 0},
	}

	res := m.pow(n)
	return res[0][2]
}
