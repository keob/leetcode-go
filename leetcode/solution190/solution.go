package solution190

func reverseBits(num uint32) (res uint32) {
	for i := 0; i < 32 && num > 0; i++ {
		res |= num & 1 << (31 - i)
		num >>= 1
	}
	return
}
