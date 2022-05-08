package solution442

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if num-1 != i {
			res = append(res, num)
		}
	}
	return res
}
