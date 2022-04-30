package solution908

func smallestRange(nums []int, k int) int {
	minNum, maxNum := nums[0], nums[0]
	for _, num := range nums[1:] {
		if num < minNum {
			minNum = num
		} else if num > maxNum {
			maxNum = num
		}
	}
	res := maxNum - minNum - 2*k
	if res > 0 {
		return res
	}
	return 0
}
