package solution372

func superPow(a int, b []int) int {
	mod := 1337
	tag := a % mod
	res := 1

	for i := len(b) - 1; i >= 0; i-- {
		dec := b[i]
		if dec != 0 {
			res = (res * pow(tag, dec, mod)) % mod
		}
		tag = pow(tag, 10, mod)
	}

	return res
}

func pow(a int, n int, mod int) int {
	tag := a % mod
	res := 1

	for n != 0 {
		if n&1 != 0 {
			res = (res * tag) % mod
		}
		tag = (tag * tag) % mod
		n >>= 1
	}

	return res
}
