package solution581

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	left := 0
	right := n - 1

	for left < n-1 && nums[left] <= nums[left+1] {
		left++
	}
	for right > 0 && nums[right] >= nums[right-1] {
		right--
	}

	if left >= right {
		return 0
	}

	min, max := getBound(nums, left, right)

	for left > 0 {
		if nums[left-1] > min {
			left--
		} else {
			break
		}
	}
	for right < n-1 {
		if nums[right+1] < max {
			right++
		} else {
			break
		}
	}

	return right - left + 1
}

func getBound(nums []int, left, right int) (min int, max int) {
	max = nums[left]
	min = nums[left]

	for left <= right {
		if nums[left] > max {
			max = nums[left]
		}
		if nums[left] < min {
			min = nums[left]
		}
		left++
	}

	return
}
