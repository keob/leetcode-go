package solution503

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	for i := range nums {
		res[i] = -1
	}

	stack := []int{}

	for i := 0; i < n*2-1; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			res[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}

	return res
}
