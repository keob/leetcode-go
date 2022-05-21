package solution961

func repeatedNTimes(nums []int) int {
	for gap := 1; gap < 3; gap++ {
		for i, num := range nums[:len(nums)-gap] {
			if num == nums[i+gap] {
				return num
			}
		}
	}
	return -1
}
