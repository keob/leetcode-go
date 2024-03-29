package offer10I

const mod int = 1e9 + 7

type matrix [2][2]int

func multiply(a matrix, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = (a[i][0]*b[0][j] + a[i][1]*b[1][j]) % mod
		}
	}
	return
}

func pow(a matrix, n int) matrix {
	res := matrix{{1, 0}, {0, 1}}

	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = multiply(res, a)
		}
		a = multiply(a, a)
	}

	return res
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	res := pow(matrix{{1, 1}, {1, 0}}, n-1)

	return res[0][0]
}
