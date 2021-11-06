package solution268

func missingNumber(nums []int) (xor int) {
	for i, num := range nums {
		xor ^= i ^ num
	}
	return xor ^ len(nums)
}
