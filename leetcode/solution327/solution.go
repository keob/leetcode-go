package solution327

func countRangeSum(nums []int, lower int, upper int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	tmp := make([]int, len(nums)+1)
	s := make([]int, len(nums)+1)

	for i := 1; i <= len(nums); i++ {
		s[i] = s[i-1] + nums[i-1]
	}

	return mergesort(s, 0, len(nums), tmp, lower, upper)
}

func mergesort(s []int, left, right int, tmp []int, lower, upper int) int {
	if left >= right {
		return 0
	}

	res := 0
	mid := (right + left) / 2
	res += mergesort(s, left, mid, tmp, lower, upper)
	res += mergesort(s, mid+1, right, tmp, lower, upper)
	l := left
	low, upp := mid+1, mid+1

	for l <= mid {
		for low <= right && s[low]-s[l] < lower {
			low++
		}
		for upp <= right && s[upp]-s[l] <= upper {
			upp++
		}
		res += upp - low
		l++
	}

	merge(s, tmp, left, mid, right)

	return res
}

func merge(s []int, tmp []int, left, mid, right int) {
	l := left
	r := mid + 1
	index := left

	for l <= mid && r <= right {
		if s[l] <= s[r] {
			tmp[index] = s[l]
			l++
			index++
		} else {
			tmp[index] = s[r]
			r++
			index++
		}
	}
	for l <= mid {
		tmp[index] = s[l]
		l++
		index++
	}
	for r <= right {
		tmp[index] = s[r]
		r++
		index++
	}
	for i := left; i <= right; i++ {
		s[i] = tmp[i]
	}
}
