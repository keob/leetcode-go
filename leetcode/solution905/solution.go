package solution905

func sortArrayByParity(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, num := range nums {
		if num%2 == 0 {
			res = append(res, num)
		}
	}
	for _, num := range nums {
		if num%2 == 1 {
			res = append(res, num)
		}
	}
	return res
}

/*
func sortArrayByParity(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	left, right := 0, n-1

	for _, num := range nums {
		if num%2 == 0 {
			res[left] = num
			left++
		} else {
			res[right] = num
			right--
		}
	}
	return res
}
*/

/*
func sortArrayByParity(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left]%2 == 0 {
			left++
		}
		for left < right && nums[right]%2 == 1 {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return nums
}
*/
