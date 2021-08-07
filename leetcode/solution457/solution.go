package solution457

func circularArrayLoop(nums []int) bool {
	n := len(nums)

	next := func(cur int) int {
		return ((cur+nums[cur])%n + n) % n
	}

	for i, num := range nums {
		if num == 0 {
			continue
		}
		slow, fast := i, next(i)
		for nums[slow]*nums[fast] > 0 && nums[slow]*nums[next(fast)] > 0 {
			if slow == fast {
				if slow == next(slow) {
					break
				}
				return true
			}
			slow = next(slow)
			fast = next(next(fast))
		}
		add := i
		for nums[add]*nums[next(add)] > 0 {
			tmp := add
			add = next(add)
			nums[tmp] = 0
		}
	}

	return false
}
